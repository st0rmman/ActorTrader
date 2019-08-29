package Quote

import (
	"../Const"
	. "../gosiris"
)

const QuoteManagerActrion  = "QuoteManagerActrion"

type QuoteManger struct {
	Base Actor
}

func NewQuoteManger()  *QuoteManger {
	qm := QuoteManger{}
	ActorSystem().RegisterActor( Const.QuoteManger , &qm.Base, nil)
	qm.Base.React("QuoteManagerActrion",qm.Action )
	return  & QuoteManger{}
}

func (actor *QuoteManger) Action(context Context) {

}