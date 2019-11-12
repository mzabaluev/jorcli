package jcli_test

import (
	"fmt"
	"strings"

	"github.com/rinor/jorcli/jcli"
)

func ExampleCertificateNewStakePoolRegistration() {
	var (
		kesKey              string
		vrfKey              string
		startValidity       = uint64(0)
		managementThreshold = uint8(1)
		serial              = uint64(1020304050)
		owner               = []string{"ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v"}
		operator            []string // no operator in this case
		outputFile          = ""     // "" - output to STDOUT only, "stakePool.cert" - will also save output to that file

		// seed used only for testing and reproducibility
		seed = "0000000000000000000000000000000000000000000000000000000000000000"
	)

	// kes25519-12-sk1qqqqqqqtawlq47yfd829sz0hjpdfs7ywhjkpy2f3z60ptpavadx9egx09zfg75v6qj9xqvea45rfpaxvu0unyt8ecvx2k7yxnlugfx980akf0qapznjg7gpfljj9pp8zqy8p2t76xksa79wxgre5t3qmtu6yme9zh2te3zfn8ctjdawxtx4umds8ljthth6f203qyah5azhpk82tsl7n3wsu97r8qssp5u36z5tdxv3rukje7hqkgmvys5vk7x4f6dds3zwcfqp7gmqxmnjv20ne5swzprjvsngtkl244v4wmw6daent3av5q0k4rqktl73qg6063ref9wa83za2e9r5t6aaqfz67hq0x5283j7f7p0fn6mtsahf5s4pukwcqdfkynzqy3fx36xj40ze6avxped0hcm5mjkdv0hznhqtrxer8u24pk92uc5nq679evw9ddpdzr8alzu5mgmth3fz3epm0k6m693tydsgutw7xjcu675pwdxavz707recqhjgc6w0264fks5r4unt862m25g5pfndnsxye53v8n5f42cmefcdvfmy930yxm8vgz69d34a8dn3v8qhcqthpukvedv63yy62mrt8anr5zwppf8g326n5h38mnzmm7nf4644x5szpgll3uhy32ejmdjmnhn8kvyt2dq8q2673yzznjgzcrruha4arzaeevzgts5xuenc9pcrdnuvku89qdyzlek72kdtszk6dmcsju3v9z5egfxsrwa3eu6tuejhqmrnjrhfmrk98ugnnz6ugayrwtwcrp7y7q98jcd2hprvavxmtesnpxn2df08vy3tk8ypwfh2a43let52tc6q3e9usylyskwyaswkt469am9f5x4v30vslutz7czayx7j4n0g52jyyx7fn57xj43z9xfarmezjew4z4q6v5fnnxh5fy6tu73xe8tremhdexuzkrlzvnq0c97jnm0d8a2uk7ewu5zqf8xaf45tag0487wf98y63ewlfdyuf53awy08heppl7nax3deahunfmr4c2sp9l8dryz7qzmc8djmxm5u6ta45dhlfx40tdwwdhm83xepn8epkw7cfw7hc27tehspcry4nhgjzkd6wdl4wn0ljnsvnwpz37ztfdy0p7yqxkzd8kvvjgwm2748mym45g2wq782ceglj0edd8lfh9g8hjhfsptdlkq2kyldk2udqukk6yzlysf0y08c7lcf9r63ngzg5cpn8kksdy85en3ljv3vl8pse2mcs60l3pyc5alke97a33vfeppn8t75d4xcp2h8u03l3gjwqj74djz09q6g4np9m0u79fupje0kd9vupa9lcs5lrthplyvet85vca96eflj7zz84806yqur4lq6se3wes8wjp9kn5xtuslk66jvx4e9vfu6tp87fw9fyekedmv9pv2579stuf4w846ql54ewdpjcd8j0gwf4ekql0t95zu80y3gc3cmzc2lxvcpt724am87pr4a3m6j6f5tc9zw52zc376ducpc6586a6vakdca6xkzm4s4m59tsvmnxya289e32v60dmkcz3pk03jxe4mrnj6wl2udfgf49eylstc6sjmr5zve0a338l8xye0n4yad5n53ka8t43kuwpycvqe4qje7nx4w4c3zpxrhxkcdvdexvl3w344cf864kekr4rgrds5luv36nn2slj73v2wxwgm4rsvf8ggazy3ytzx6dwvc8986wretny8mua6maf3fdtrdyp0c787d8tjg30xkvwfyll7s4pj0yn8s3lgyg5qsc6v5w0sa0gs0l5s7r7mnn0q6rvthkvkag226yw0xfmake7kjsrsv4klu733fpwrk34lq3zphrzamcge60z425x34l3g7qj49l4d53rfrl6emvr9dxagflldu32550c4uz4wrvxq0fa463cazxvhnl9
	kesPrivateKey, err := jcli.KeyGenerate(seed, "SumEd25519_12", "")
	if err != nil {
		fmt.Printf("kesPrivateKey FAILED: %s\n", err)
		return
	}

	// kes25519-12-pk1j4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4s4zrqja
	kesPublicKey, err := jcli.KeyToPublic(kesPrivateKey, "", "")
	if err != nil {
		fmt.Printf("kesPublicKey FAILED: %s\n", err)
		return
	}

	// vrf_sk1fffu87autxtsae0cttup8p6allqn4yz29ef6uln9lgx75mnzeyqshc9lsp
	vrfPrivateKey, err := jcli.KeyGenerate(seed, "Curve25519_2HashDH", "")
	if err != nil {
		fmt.Printf("vrfPrivateKey FAILED: %s\n", err)
		return
	}

	// vrf_pk1n3n2xwwgx38eyt7rypkttkhgzjjefsqh0hfjxhp9fkwypxn9hqyq6870lk
	vrfPublicKey, err := jcli.KeyToPublic(vrfPrivateKey, "", "")
	if err != nil {
		fmt.Printf("vrfPublicKey FAILED: %s\n", err)
		return
	}

	// convert to string from []byte and remove newline/space
	kesKey = strings.TrimSpace(string(kesPublicKey))
	vrfKey = strings.TrimSpace(string(vrfPublicKey))

	stakePoolCert, err := jcli.CertificateNewStakePoolRegistration(kesKey, vrfKey, startValidity, managementThreshold, serial, owner, operator, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakePoolRegistration: %s", err)
	} else {
		fmt.Printf("%s", stakePoolCert)
	}
	// Output:
	//
	// cert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqa8dlhx
}

func ExampleCertificateSign_registration_stdin() {
	var (
		stdinCert      = []byte("cert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqa8dlhx")
		signingKeyFile = []string{"testdata/private_key_txt.golden"} // ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
		inputFile      = ""                                          // "" - input from STDIN (stdinCert []byte), "stakePool.cert" - will load the certificate from that file
		outputFile     = ""                                          // "" - output to STDOUT ([]byte) only, "stakePool.signed_cert" - will also save output to that file
	)

	stakePoolSignedCert, err := jcli.CertificateSign(stdinCert, signingKeyFile, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateSign: %s", err)
	} else {
		fmt.Printf("%s", stakePoolSignedCert)
	}
	// Output:
	//
	// signedcert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpqp74fv2vt93cvctyjp6fnsmhnxt732g3tvsxk0pdrym34ecwhwu2ua9hpjqggntpv7fyw3k4ynf0epc04v7aut0flpn6yddvfy343qgqyzh80g
}

func ExampleCertificateGetStakePoolID_stdin() {
	var (
		stdinCertSigned = []byte("cert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqa8dlhx")
		inputFile       = "" // "" - input from STDIN (stdinCertSigned []byte), "stakePool.signed_cert" - will load the certificate from that file
		outputFile      = "" // "" - output to STDOUT ([]byte) only, "stakePool.id" - will also save output to that file
	)

	stakePoolID, err := jcli.CertificateGetStakePoolID(stdinCertSigned, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateGetStakePoolID: %s", err)
	} else {
		fmt.Printf("%s", stakePoolID)
	}
	// Output:
	//
	// 205b0d6d0ddf3a5cc9f716cc7915a43dea3b2647a8b4efca5cbbd3489c0708c5
}

func ExampleCertificateNewStakeDelegation() {
	var (
		stakePoolID = "205b0d6d0ddf3a5cc9f716cc7915a43dea3b2647a8b4efca5cbbd3489c0708c5"
		stakeKey    = "ed25519_pk10p43s2c5g3hhdklz9k6awwy5nvv7cnkwv6szgaxvac4ju0jm2a0qyf6j8v" // Public Key
		outputFile  = ""                                                                      // "" - output to STDOUT ([]byte) only, "stakePoolDelegation.new" - will also save output to that file
	)

	stakeNewDeleg, err := jcli.CertificateNewStakeDelegation(stakeKey, []string{stakePoolID}, outputFile)

	if err != nil {
		fmt.Printf("CertificateNewStakeDelegation: %s", err)
	} else {
		fmt.Printf("%s", stakeNewDeleg)
	}
	// Output:
	//
	// cert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqfqtvxk6rwl8fwvnacke3u3tfpaagajv3agknhu5h9m6dyfcpcgc5r6uzuv
}

func ExampleCertificateSign_delegation_stdin() {
	var (
		stdinCert      = []byte("cert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqfqtvxk6rwl8fwvnacke3u3tfpaagajv3agknhu5h9m6dyfcpcgc5r6uzuv")
		signingKeyFile = []string{"testdata/private_key_txt.golden"} // ed25519e_sk1wzuwptdq7y7eqszadtj48p4a9z7ayxdc5zx76x4gxmhuezmhp4ra5s2e03g4wjydwujwq0acmp9rw6jrhr6p2x9prnpc0dnfkthxtps9029w4
		inputFile      = ""                                          // "" - input from STDIN (stdinCert []byte), "stakePool.cert" - will load the certificate from that file
		outputFile     = ""                                          // "" - output to STDOUT ([]byte) only, "stakePool.signed_cert" - will also save output to that file
	)

	stakePoolSignedCert, err := jcli.CertificateSign(stdinCert, signingKeyFile, inputFile, outputFile)

	if err != nil {
		fmt.Printf("CertificateSign: %s", err)
	} else {
		fmt.Printf("%s", stakePoolSignedCert)
	}
	// Output:
	//
	// signedcert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqfqtvxk6rwl8fwvnacke3u3tfpaagajv3agknhu5h9m6dyfcpcgc5q4qy45fktgfrt26atr4raujxqmztdnsfg0rkl6fgyecr8xg8820l9caa62j9yzq4skppm8y5ly8n8gkckems0xda8t69q2cvv4j2z2qu7764rq
}

func ExampleCertificatePrint_registrationSigned_stdin() {
	var (
		stdinCert = []byte("signedcert1qvqqqqqqqqqqqqqqqqqqq0xsn2eqqqqqqqqqqqqqqqqqqqqqqqqqr8rx5vuusdz0jghuxgrvkhdws999jnqpwlwnydwz2nvugzdxtwqgj4re3s6yqj7js55w03pyrtguvwvfm2cezv8qy7xcfvrf5yg3fc4sz7rtrq43g3r0wmd7ytd46uuffxcea38vue4qy36vem3t9cl9k467qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqpqp74fv2vt93cvctyjp6fnsmhnxt732g3tvsxk0pdrym34ecwhwu2ua9hpjqggntpv7fyw3k4ynf0epc04v7aut0flpn6yddvfy343qgqyzh80g")
		inputFile = "" // "" - input from STDIN (stdinCert []byte), "stakePool.signed_cert" - will load the certificate from that file
	)

	certPrint, err := jcli.CertificatePrint(stdinCert, inputFile)

	if err != nil {
		fmt.Printf("CertificatePrint: %s", err)
	} else {
		fmt.Printf("%s", certPrint)
	}
	// Output:
	//
	// Certificate(PoolRegistration(PoolRegistration { serial: 1020304050, start_validity: TimeOffsetSeconds(DurationSeconds(0)), permissions: PoolPermissions(1), owners: [786b182b14446f76dbe22db5d738949b19ec4ece66a02474ccee2b2e3e5b575e], operators: [], rewards: TaxType { fixed: Value(0), ratio: Ratio { numerator: 0, denominator: 1 }, max_limit: None }, reward_account: None, keys: GenesisPraosLeader { kes_public_key: 954798c34404bd28528e7c4241ad1c63989dab19130e0278d84b069a11114e2b, vrf_public_key: 9c66a339c8344f922fc3206cb5dae814a594c0177dd3235c254d9c409a65b808 } }))
}

func ExampleCertificatePrint_delegationSigned_stdin() {
	var (
		stdinCert = []byte("signedcert1q9uxkxptz3zx7akmugkmt4ecjjd3nmzween2qfr5enhzkt37tdt4uqfqtvxk6rwl8fwvnacke3u3tfpaagajv3agknhu5h9m6dyfcpcgc5q4qy45fktgfrt26atr4raujxqmztdnsfg0rkl6fgyecr8xg8820l9caa62j9yzq4skppm8y5ly8n8gkckems0xda8t69q2cvv4j2z2qu7764rq")
		inputFile = "" // "" - input from STDIN (stdinCert []byte), "stakePoolDelegation.new" - will load the certificate from that file
	)

	certPrint, err := jcli.CertificatePrint(stdinCert, inputFile)

	if err != nil {
		fmt.Printf("CertificatePrint: %s", err)
	} else {
		fmt.Printf("%s", certPrint)
	}
	// Output:
	//
	// Certificate(StakeDelegation(StakeDelegation { account_id: UnspecifiedAccountIdentifier([120, 107, 24, 43, 20, 68, 111, 118, 219, 226, 45, 181, 215, 56, 148, 155, 25, 236, 78, 206, 102, 160, 36, 116, 204, 238, 43, 46, 62, 91, 87, 94]), delegation: Full($ hash_ty(0x205b0d6d0ddf3a5cc9f716cc7915a43dea3b2647a8b4efca5cbbd3489c0708c5)) }))
}
