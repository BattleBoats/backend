package models

type GameBoard struct {
	// Board *string `json:"data,omitempty"`
	// PlayerOne *GamePlayer `json:"player_one,omitempty"`
	// PlayerTwo *GamePlayer `json:"player_two,omitempty"`
}

type GamePlayer struct {
	Id    *int64  `json:"id,omitempty"`
	Ships *[]Ship `json:"ships,omitempty"`
}

/*
{
    "player_one" : {
        "id" : 3,
        "ships" : [
            {
                "size" : 5,
                "hits" : [true, false, false, true, false],
                "origin" : {
                    "x" : 14,
                    "y" : 8
                },
                "orientation" : 2,
                "is_using_special" : true
            },
            {
                "size" : 2,
                "hits" : [false, true],
                "origin" : {
                    "x" : 0,
                    "y" : 0
                },
                "orientation" : 0,
                "is_using_special" : false
            }
        ]
    },
    "player_two" : {

    }
}
*/
