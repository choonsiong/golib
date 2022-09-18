package file

import "fmt"

// convertToBinary returns the permissions given in binary form.
func convertToBinary(permissions string) (string, error) {
	if permissions == "" {
		return "", fmt.Errorf("%w: %q", ErrInvalidTriplet, permissions)
	}

	binaryPermissions := permissions[1:]

	p1, err := tripletToBinary(binaryPermissions[0:3])
	if err != nil {
		return "", err
	}

	p2, err := tripletToBinary(binaryPermissions[3:6])
	if err != nil {
		return "", err
	}

	p3, err := tripletToBinary(binaryPermissions[6:9])
	if err != nil {
		return "", err
	}

	return p1 + p2 + p3, nil
}

// tripletToBinary returns the single triplet in three binary digits.
func tripletToBinary(triplet string) (string, error) {
	if triplet == "rwx" {
		return "111", nil
	}

	if triplet == "-wx" {
		return "011", nil
	}

	if triplet == "--x" {
		return "001", nil
	}

	if triplet == "---" {
		return "000", nil
	}

	if triplet == "r-x" {
		return "101", nil
	}

	if triplet == "r--" {
		return "100", nil
	}

	if triplet == "rw-" {
		return "110", nil
	}

	if triplet == "-w-" {
		return "010", nil
	}

	return "", fmt.Errorf("%w: %q", ErrInvalidTriplet, triplet)
}
