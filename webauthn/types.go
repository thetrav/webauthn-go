package webauthn

import (
	
    "github.com/google/uuid"
)

func regOpts(challenge []byte, userName string) RegistrationOptions {
    return RegistrationOptions {
        Challenge: challenge,
        Rp: Rp {
            Name: "Local Test Server",
            Id: "localhost",
        },
        User: User {
            Id: []byte(uuid.NewString()),
            Name: userName,
            DisplayName: userName,
        },
        PubKeyCredParams: []PubKeyCredParam{PubKeyCredParam{Alg: -7, Type: "public-key"}},
    }
}

type Rp struct {
    Name string `json:"name"`
    Id string `json:"id"`
}

type User struct {
    Id []byte `json:"id"`
    Name string `json:"name"`
    DisplayName string `json:"displayName"`
}

type PubKeyCredParam struct {
    Alg int  `json:"alg"`
    Type string `json:"type"`
}

type RegistrationOptions struct {
    Challenge []byte `json:"challenge"`
    Rp Rp `json:"rp"`
    User User `json:"user"`
    PubKeyCredParams []PubKeyCredParam `json:"pubKeyCredParams"`
}

type Credential struct {
    //TODO: define credential payload
}

func MakeChallenge() byte[], error {
    challenge := make([]byte, 32)
    if n, err := rand.Read(challenge); err != nil {
        return nil, err
    } else if n != len(challenge) {
        return nil, fmt.Errorf("tried to generate 32 random bytes, but got %d\n", n)
    }
    return challenge, nil
}

