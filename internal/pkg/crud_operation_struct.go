package specs

import (
	"fmt"
	"regexp"

	"github.com/purisaurabh/database-connection/internal/pkg/constants"
	"github.com/purisaurabh/database-connection/internal/pkg/errors"
)

type PostProfile struct {
	Profiles PostRequest `json:"profiles"`
}

type PostRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type ListProfileResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Created_At int64  `json:"created_at"`
	Updated_At int64  `json:"updated_at"`
}

type UpdateProfile struct {
	Profiles UpdateRequest `json:"profiles"`
}

type UpdateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Updated_at int64  `json:"updated_at"`
}

func (p *PostProfile) Validate() error {
	if p.Profiles.Name == "" {
		return errors.ErrParameterMissing
	}
	if p.Profiles.Email == "" {
		return errors.ErrParameterMissing
	}
	if p.Profiles.Mobile == "" {
		return errors.ErrParameterMissing
	}

	matchMail, err := regexp.MatchString(constants.EmailRegex, p.Profiles.Email)
	if err != nil {
		fmt.Println("Error in matching email regex", err)
		return err
	}
	if !matchMail {
		return fmt.Errorf("%s : email ", errors.ErrInvalidFormat.Error())
	}

	matchMobile, err := regexp.MatchString(constants.MobileRegex, p.Profiles.Mobile)
	if err != nil {
		fmt.Println("Error in matching mobile regex", err)
		return err
	}
	if !matchMobile {
		return fmt.Errorf("%s : mobile ", errors.ErrInvalidFormat.Error())
	}

	return nil
}

func (u *UpdateProfile) Validate() error {
	if u.Profiles.Name == "" {
		return errors.ErrParameterMissing
	}
	if u.Profiles.Email == "" {
		return errors.ErrParameterMissing
	}
	if u.Profiles.Mobile == "" {
		return errors.ErrParameterMissing
	}

	matchMail, err := regexp.MatchString(constants.EmailRegex, u.Profiles.Email)
	if err != nil {
		fmt.Println("Error in matching email regex", err)
		return err
	}
	if !matchMail {
		return fmt.Errorf("%s : email ", errors.ErrInvalidFormat.Error())
	}

	matchMobile, err := regexp.MatchString(constants.MobileRegex, u.Profiles.Mobile)
	if err != nil {
		fmt.Println("Error in matching mobile regex", err)
		return err
	}
	if !matchMobile {
		return fmt.Errorf("%s : mobile ", errors.ErrInvalidFormat.Error())
	}

	return nil
}
