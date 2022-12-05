package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerPerformanceIndicatorsResponseEntity struct{
    Type	string	`json:"type"`
    Name	string	`json:"name"`
    Tip	string	`json:"tip"`
    Score	int	`json:"score"`
    ScoreFormat	string	`json:"score_format"`
    FormattedScore	string	`json:"formatted_score"`
    Target	int	`json:"target"`
    TargetFormat	string	`json:"target_format"`
    FormattedTarget	string	`json:"formatted_target"`
    TargetRespected	bool	`json:"target_respected"`
    ActionUrl	string	`json:"action_url"`
}
func (g GetSellerPerformanceIndicatorsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
