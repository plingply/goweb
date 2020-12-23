package card

import (
	"errors"
	"goframe-web/app/model"
	"unicode/utf8"
)

func GetCardList(school_id, campus_id, page, limit uint) (result interface{}, total int, err error) {

	if school_id == 0 {
		return nil, 0, errors.New("学校id不能为空")
	}

	if campus_id == 0 {
		return nil, 0, errors.New("校区id不能为空")
	}

	if page <= 0 {
		return nil, 0, errors.New("参数错误")
	}
	if limit <= 0 {
		return nil, 0, errors.New("参数错误")
	}

	var card *model.Card
	result, total = card.GetCardList(school_id, campus_id, page, limit)

	return result, total, nil
}

func GetCardSimpleList(school_id, campus_id uint) (result interface{}, err error) {
	if school_id == 0 {
		return nil, errors.New("学校id不能为空")
	}
	var card *model.Card
	result = card.GetCardSimpleList(school_id, campus_id)
	return result, nil
}

func UpdateCard(card_id uint, data map[string]interface{}) (re bool, msg error) {

	if card_id <= 0 {
		return false, errors.New("参数错误")
	}

	if data["card_name"] != nil && utf8.RuneCountInString(data["card_name"].(string)) > 20 {
		return false, errors.New("参数错误 card_name")
	}

	if data["remark"] != nil && utf8.RuneCountInString(data["remark"].(string)) > 100 {
		return false, errors.New("参数错误 remark")
	}

	var card *model.Card
	re = card.UpdateCard(card_id, data)

	return
}

func CreateCard(school_id, campus_id uint, card *model.Card) (re bool, msg error) {

	if school_id <= 0 {
		return false, errors.New("参数错误")
	}

	if campus_id <= 0 {
		return false, errors.New("参数错误")
	}

	var cardModel *model.Card
	re = cardModel.CreateCard(card)

	return
}
