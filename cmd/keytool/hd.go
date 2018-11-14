package main

import (
	"errors"
	"fmt"

	cli "github.com/jawher/mow.cli"
	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/oneiro-ndev/ndaumath/pkg/key"
)

func hdstr(k key.ExtendedKey) string {
	text, err := k.MarshalText()
	check(err)
	return string(text)
}

func hdparse(s string) *key.ExtendedKey {
	k := new(key.ExtendedKey)
	err := k.UnmarshalText([]byte(s))
	check(err)
	return k
}

func cmdHDNew(cmd *cli.Cmd) {
	cmd.Action = func() {
		seed, err := key.GenerateSeed(key.RecommendedSeedLen)
		check(err)
		k, err := key.NewMaster(seed)
		check(err)
		fmt.Println(hdstr(*k))
	}
}

func cmdHDPublic(cmd *cli.Cmd) {
	cmd.Spec = getKeySpec("PVT")
	getKey := getKeyClosureHD(cmd, "PVT", "private key from which to make a public key")

	cmd.Action = func() {
		pvt := getKey()
		pub, err := pvt.Public()
		check(err)
		fmt.Println(hdstr(*pub))
	}
}

func cmdHDChild(cmd *cli.Cmd) {
	cmd.Spec = fmt.Sprintf(
		"%s PATH",
		getKeySpec(""),
	)

	getKey := getKeyClosureHD(cmd, "", "key from which to derive a child")

	pathS := cmd.StringArg("PATH", "", "derivation path for child key")
	cmd.Action = func() {
		key := getKey()
		key, err := key.DeriveFrom("/", *pathS)
		check(err)
		fmt.Println(hdstr(*key))
	}
}

func cmdHDConvert(cmd *cli.Cmd) {
	keyS := cmd.StringArg("KEY", "", "old-format key to convert")

	cmd.Action = func() {
		k, err := key.FromOldSerialization(*keyS)
		check(err)
		fmt.Println(hdstr(*k))
	}
}

func cmdHDTruncate(cmd *cli.Cmd) {
	cmd.Spec = getKeySpec("")

	getKey := getKeyClosureHD(cmd, "", "key to truncate")

	cmd.Action = func() {
		key := getKey()
		var keyB []byte
		var err error
		if key.IsPrivate() {
			skey, err := key.SPrivKey()
			check(err)
			skey.Truncate()
			keyB, err = skey.MarshalText()
		} else {
			skey, err := key.SPubKey()
			check(err)
			skey.Truncate()
			keyB, err = skey.MarshalText()
		}
		check(err)
		fmt.Println(string(keyB))
	}
}

func cmdHDAddr(cmd *cli.Cmd) {
	// mow.cli ensures with this that only one option is specified
	cmd.Spec = fmt.Sprintf(
		"%s [-k=<kind> | -a | -n | -e | -x]",
		getKeySpec(""),
	)

	getKey := getKeyClosureHD(cmd, "", "get address from this key, converting to public as necessary")

	var (
		pkind  = cmd.StringOpt("k kind", string(address.KindUser), "manually specify address kind")
		kuser  = cmd.BoolOpt("a user", false, "address kind: user (default)")
		kndau  = cmd.BoolOpt("n ndau", false, "address kind: ndau")
		kendow = cmd.BoolOpt("e endowment", false, "address kind: endowment")
		kxchng = cmd.BoolOpt("x exchange", false, "address kind: exchange")
	)

	cmd.Action = func() {
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

		key := getKey()
		addr, err := address.Generate(kind, key.PubKeyBytes())
		check(err)
		fmt.Println(addr)
	}
}
