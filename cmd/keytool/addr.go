package main

import (
	"errors"
	"fmt"

	cli "github.com/jawher/mow.cli"
	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/oneiro-ndev/ndaumath/pkg/signature"
)

func getKindSpec() string {
	// mow.cli ensures with this that only one option is specified
	return "[-k=<kind> | -a | -n | -e | -x]"
}

func getKindClosure(cmd *cli.Cmd) func() address.Kind {
	var (
		pkind  = cmd.StringOpt("k kind", string(address.KindUser), "manually specify address kind")
		kuser  = cmd.BoolOpt("a user", false, "address kind: user (default)")
		kndau  = cmd.BoolOpt("n ndau", false, "address kind: ndau")
		kendow = cmd.BoolOpt("e endowment", false, "address kind: endowment")
		kxchng = cmd.BoolOpt("x exchange", false, "address kind: exchange")
	)

	return func() address.Kind {
		kind := address.Kind(*pkind) // never nil dereference; defaults to user
		if kuser != nil && *kuser {
			kind = address.KindUser
		}
		if kndau != nil && *kndau {
			kind = address.KindNdau
		}
		if kendow != nil && *kendow {
			kind = address.KindEndowment
		}
		if kxchng != nil && *kxchng {
			kind = address.KindExchange
		}

		if !address.IsValidKind(kind) {
			check(errors.New("invalid kind: " + string(kind)))
		}
		return kind
	}
}

func cmdAddr(cmd *cli.Cmd) {
	cmd.Spec = fmt.Sprintf(
		"%s %s",
		getKeySpec("PUB"),
		getKindSpec(),
	)

	getKey := getKeyClosure(cmd, "PUB", "get address from this key")
	getKind := getKindClosure(cmd)

	cmd.Action = func() {
		key := getKey()
		_, ok := key.(*signature.PublicKey)
		if !ok {
			check(errors.New("addresses can only be generated from public keys"))
		}

		kind := getKind()

		addr, err := address.Generate(kind, key.KeyBytes())
		check(err)
		fmt.Println(addr)
	}
}
