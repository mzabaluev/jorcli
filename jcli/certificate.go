package jcli

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// CertificateGetStakePoolID - get the stake pool id from the given stake pool registration certificate.
//
//  [STDIN] | jcli certificate get-stake-pool-id [<FILE_INPUT>] [<FILE_OUTPUT>] | [STDOUT]
func CertificateGetStakePoolID(
	stdinCertSigned []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCertSigned) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCertSigned", "inputFile")
	}

	arg := []string{"certificate", "get-stake-pool-id"}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCertSigned = nil
	}
	if outputFile != "" && inputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := jcli(stdinCertSigned, arg...)
	if err != nil /* || outputFile == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	//
	// convert stdout to outputFile
	if outputFile != "" && inputFile == "" {
		if err = ioutil.WriteFile(outputFile, out, 0644); err != nil {
			return nil, err
		}
	}
	if outputFile == "" {
		return out, nil
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewOwnerStakeDelegation - build an owner stake delegation certificate.
//
//  jcli certificate new owner-stake-delegation <STAKE_POOL_ID:weight>... [--output <output>] | [STDOUT]
func CertificateNewOwnerStakeDelegation(
	weightedPoolID []string,
	outputFile string,
) ([]byte, error) {
	if len(weightedPoolID) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "weightedPoolID")
	}

	maxPools := 8 // The maximum number of delegation pools
	if len(weightedPoolID) > maxPools {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "weightedPoolID", 1, maxPools, len(weightedPoolID))
	}

	arg := []string{
		"certificate", "new", "owner-stake-delegation",
	}
	arg = append(arg, weightedPoolID...) // FIXME: should check data validity!

	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakeDelegation - build a stake delegation certificate.
//
//  jcli certificate new stake-delegation <STAKE_KEY> <STAKE_POOL_ID:weight>... [--output <output>] | [STDOUT]
func CertificateNewStakeDelegation(
	stakeKey string,
	weightedPoolID []string,
	outputFile string,
) ([]byte, error) {
	if stakeKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "stakeKey")
	}
	if len(weightedPoolID) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "weightedPoolID")
	}

	maxPools := 8 // The maximum number of delegation pools
	if len(weightedPoolID) > maxPools {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "weightedPoolID", 1, maxPools, len(weightedPoolID))
	}

	arg := []string{
		"certificate", "new", "stake-delegation",
		stakeKey,
	}
	arg = append(arg, weightedPoolID...) // FIXME: should check data validity!

	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakePoolRegistration - build a stake pool registration certificate with single/multiple owners.
//
//  jcli certificate new stake-pool-registration
//                                              --kes-key <KES_KEY>
//                                              --vrf-key <VRF_KEY>
//                                              --start-validity <SECONDS-SINCE-START>
//                                              --management-threshold <THRESHOLD>
//                                              --owner <OWNER_PUBLIC_KEY> --owner... (max owners TBD)
//                                              [--operator <OPERATOR_PUBLIC_KEY> --operator... (max operators TBD)]
//                                              [--tax-fixed <TAX_VALUE>]
//                                              [--tax-ratio <TAX_RATIO>]
//                                              [--tax-limit <TAX_LIMIT>]
//                                              [--reward-account <REWARD_ACCOUNT>]
//                                              [output] | STDOUT
func CertificateNewStakePoolRegistration(
	kesKey string,
	vrfKey string,
	startValidity uint64,
	managementThreshold uint8,
	owner []string,
	operator []string,
	taxFixed uint64,
	taxRatio string,
	taxLimit uint64,
	rewardAccount string,
	outputFile string,
) ([]byte, error) {
	if kesKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "kesKey")
	}
	if vrfKey == "" {
		return nil, fmt.Errorf("parameter missing : %s", "vrfKey")
	}
	if len(owner) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "owner")
	}
	// TODO: Confirm/Fix the limits
	/*

		maxOwners := 31   // 5 bits for the owners for a maximum of 31 elements
		maxOperators := 3 // 2 bits for the operators for a maximum of 3 elements
		if len(owner) > maxOwners {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "owner", 1, maxOwners, len(owner))
		}
		if len(operator) > maxOperators {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "operator", 0, maxOperators, len(operator))
		}

		// managementThreshold <= #owners and > 0
		if managementThreshold < 1 || int(managementThreshold) > len(owner) {
			return nil, fmt.Errorf("%s expected between %d - %d, got %d", "managementThreshold", 1, len(owner), managementThreshold)
		}
	*/
	arg := []string{
		"certificate", "new", "stake-pool-registration",
		"--kes-key", kesKey,
		"--vrf-key", vrfKey,
		"--start-validity", strconv.FormatUint(startValidity, 10),
		"--management-threshold", strconv.FormatUint(uint64(managementThreshold), 10),
	}
	for _, ownerPublicKey := range owner {
		arg = append(arg, "--owner", ownerPublicKey) // FIXME: should check data validity!
	}
	for _, operatorPublicKey := range operator {
		arg = append(arg, "--operator", operatorPublicKey) // FIXME: should check data validity!
	}

	if taxFixed > 0 {
		arg = append(arg, "--tax-fixed", strconv.FormatUint(taxFixed, 10))
	}
	if taxRatio != "" {
		arg = append(arg, "--tax-ration", taxRatio)
	}
	if taxLimit > 0 {
		arg = append(arg, "--tax-limit", strconv.FormatUint(taxLimit, 10))
	}
	if rewardAccount != "" {
		arg = append(arg, "--reward-account", rewardAccount)
	}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewStakePoolRetirement - retire the given stake pool ID From the blockchain.
// By doing so all remaining stake delegated to this stake pool will become pending and will need to be re-delegated.
//
// jcli certificate new stake-pool-retirement --pool-id <POOL_ID> --retirement-time <SECONDS-SINCE-START> [output]  | STDOUT
func CertificateNewStakePoolRetirement(
	poolID string,
	retirementTime uint64,
	outputFile string,
) ([]byte, error) {
	if poolID == "" {
		return nil, fmt.Errorf("parameter missing : %s", "poolID")
	}

	arg := []string{
		"certificate", "new", "stake-pool-retirement",
		"--pool-id", poolID,
		"--retirement-time", strconv.FormatUint(retirementTime, 10),
	}
	if outputFile != "" {
		arg = append(arg, outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateNewVotePlan - create a vote plan certificate form given config data/file.
//
//  STDIN | jcli certificate new vote-plan [<FILE_INPUT>] [--output <FILE_OUTPUT>] | [STDOUT]
func CertificateNewVotePlan(
	stdinConfig []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinConfig) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinConfig", "inputFile")
	}

	arg := []string{
		"certificate", "new", "vote-plan",
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinConfig, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateGetVotePlanID - get the vote plan id from the given vote plan certificate.
//
//  [STDIN] | jcli certificate get-vote-plan-id [<FILE_INPUT>] [<FILE_OUTPUT>] | [STDOUT]
func CertificateGetVotePlanID(
	stdinCert []byte,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}

	arg := []string{"certificate", "get-vote-plan-id"}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCert = nil
	}
	if outputFile != "" && inputFile != "" {
		arg = append(arg, outputFile) // TODO: UPSTREAM unify with "--output" as other file output commands
	}

	out, err := jcli(stdinCert, arg...)
	if err != nil /* || outputFile == "" */ {
		return out, err
	}

	// TODO: Remove this once/if UPSTREAM fixed (--input and --output)
	//
	// convert stdout to outputFile
	if outputFile != "" && inputFile == "" {
		if err = ioutil.WriteFile(outputFile, out, 0644); err != nil {
			return nil, err
		}
	}
	if outputFile == "" {
		return out, nil
	}

	return ioutil.ReadFile(outputFile)
}

// CertificateSign - Sign certificate,
// you can call this command multiple time to add multiple signatures if this is required.
//
//  [STDIN] | jcli certificate sign --key=<signing-key file>... [--certificate=<input file>] [--output=<output file>] | [STDOUT]
func CertificateSign(
	stdinCert []byte,
	signingKeyFile []string,
	inputFile string,
	outputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}
	if len(signingKeyFile) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "signingKeyFile")
	}

	arg := []string{"certificate", "sign"}
	for _, signKeyFile := range signingKeyFile {
		arg = append(arg, "--key", signKeyFile) // FIXME: should check data validity!
	}
	if inputFile != "" {
		arg = append(arg, "--certificate", inputFile)
		stdinCert = nil
	}
	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(stdinCert, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}

// CertificatePrint - Print certificate.
//
//  [STDIN] | jcli certificate print [<input file>] | STDOUT
func CertificatePrint(
	stdinCert []byte,
	inputFile string,
) ([]byte, error) {
	if len(stdinCert) == 0 && inputFile == "" {
		return nil, fmt.Errorf("%s : EMPTY and parameter missing : %s", "stdinCert", "inputFile")
	}

	arg := []string{"certificate", "print"}
	if inputFile != "" {
		arg = append(arg, inputFile) // TODO: UPSTREAM unify with "--input" as other file input commands
		stdinCert = nil
	}

	return jcli(stdinCert, arg...)
}

/*
// CertificateNewVotePlan - create a vote plan certificate.
//
//  jcli certificate new vote-plan --vote-start <vote-start> --vote-end <vote-end> --committee-end <committee-end> [--proposal-id <proposals>...] [--output <output>] | [STDOUT]
func CertificateNewVotePlanOLD(
	voteStart string,
	voteEnd string,
	committeeEnd string,
	proposalID []string,
	outputFile string,
) ([]byte, error) {
	if voteStart == "" {
		return nil, fmt.Errorf("parameter missing : %s", "voteStart")
	}
	if voteEnd == "" {
		return nil, fmt.Errorf("parameter missing : %s", "voteEnd")
	}
	if committeeEnd == "" {
		return nil, fmt.Errorf("parameter missing : %s", "committeeEnd")
	}

	if len(proposalID) == 0 {
		return nil, fmt.Errorf("parameter missing : %s", "proposalID")
	}

	maxProposals := 255 // The maximum number of proposals per voteplan
	if len(proposalID) > maxProposals {
		return nil, fmt.Errorf("%s expected between %d - %d, got %d", "proposalID", 1, maxProposals, len(proposalID))
	}

	arg := []string{
		"certificate", "new", "vote-plan",
		"--vote-start", voteStart,
		"--vote-end", voteEnd,
		"--committee-end", committeeEnd,
	}

	for _, proposal := range proposalID {
		arg = append(arg, "--proposal-id", proposal)
	}

	if outputFile != "" {
		arg = append(arg, "--output", outputFile)
	}

	out, err := jcli(nil, arg...)
	if err != nil || outputFile == "" {
		return out, err
	}

	return ioutil.ReadFile(outputFile)
}
*/
