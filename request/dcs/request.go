package dcs

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"MyCodeArchive_Go/utils/strings_"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func CreateRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"Name": "Relation4",
	}, CreateRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(RelationParam)

	err = checkNameParam(params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, false)
	if err != nil {
		return err
	}

	for _, strategyUUID := range params.StrategyIds {
		err = checkExisted("", strategyUUID, StrategyModule, true)
		if err != nil {
			return err
		}
	}

	// init relation to db
	initRelation := DcsRelations{
		UUID:  uuid.NewString(),
		State: Busy,
	}
	if err = initRelation.Create(); err != nil {
		return err
	}
	params.UUID = initRelation.UUID
	logging.Log.Infof("relation init successfully, UUID: %s", initRelation.UUID)

	updateRelation := DcsRelations{
		UUID: initRelation.UUID,
	}
	if err = updateRelation.Update(map[string]interface{}{
		"name":               params.Name,
		"master_pool":        params.MasterPool,
		"slave_pool":         params.SlavePool,
		"master_resource_id": params.MasterResource.ID,
		"slave_resource_id":  params.SlaveResource.ID,
		"resource_type":      params.ResourceType,
		"strategy_ids":       strings.Join(params.StrategyIds, ","),
		"last_sync_time":     params.LastSyncTime,
		"last_sync_snap":     params.LastSyncSnap,
		"state":              Idle,
		"running_state":      Nomal,
		"health_state":       params.HealthState,
		"data_state":         Incomplete,
		"role":               params.Role,
		"is_config_sync":     params.IsConfigSync,
	}); err != nil {
		logging.Log.Infof("rollback uuid:%s", params.UUID)
		rollbackErr := deleteRelation2DBRollback(params)
		if rollbackErr != nil {
			logging.Log.Errorf("rollback failed. %+v", params)
		}
		return err
	}

	fmt.Println("end craete raltions")
	return nil
}

func UpdateRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID":    "dcba7a69-4991-45d4-953c-47fbf10a7f66",
		"Name":    "",
		"NewName": "relationNewName3",
	}, UpdateRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(RelationUpdateParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkNameParam(params.NewName)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	err = checkExisted(params.NewName, "", RelationModule, false)
	if err != nil {
		return err
	}

	var relation DcsRelations
	relation.Name = params.Name
	relation.UUID = params.UUID
	update := map[string]interface{}{
		"name": params.NewName,
	}
	if err = relation.Update(update); err != nil {
		return err
	}
	fmt.Println("end update raltions")
	return nil
}

func DeleteRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relationNewName3",
	}, DeleteRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	var relation DcsRelations
	relation.Name = params.Name
	relation.UUID = params.UUID
	if err = relation.Delete(); err != nil {
		return err
	}
	fmt.Println("end delete raltions")
	return nil
}

func ShowRelationExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "",
		"Name": "relation3",
	}, ShowRelationFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, RelationModule, true)
	if err != nil {
		return err
	}

	relation := DcsRelations{UUID: params.UUID, Name: params.Name}
	if len(params.UUID) != 0 {
		err = relation.QueryById()
	} else {
		err = relation.QueryByName()
	}
	if err != nil {
		return err
	}
	logging.Log.Infof("result: %+v", relation)
	return nil
}

func ListRelationsExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{}, ListRelationsFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	err = CheckSearchParam(params.SearchParam)
	if err != nil {
		return err
	}

	var relationDb DcsRelations
	relations, total, err := relationDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	logging.Log.Infof("%+v", relations)
	logging.Log.Infof("lenth: %d", total)
	return nil
}

func ListRelationsStrategiesExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{}, ListRelationsStrategiesFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	err = CheckSearchParam(params.SearchParam)
	if err != nil {
		return err
	}

	var relationDb DcsRelations
	relations, total, err := relationDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	var strategyDb DcsStrategies
	relationsStrategies := []RelationsStrategies{}
	for _, relation := range relations {
		strategies, err := strategyDb.QueryByIds(strings.Split(relation.StrategyIds, ","))
		if err != nil {
			return err
		}
		for index, strategy := range strategies {
			strategies[index].Description = strings_.UnquoteString(strategy.Description)
		}
		relationsStrategies = append(relationsStrategies, RelationsStrategies{
			DcsRelations: relation,
			Strategies:   strategies,
		})
	}

	logging.Log.Infof("%+v", relationsStrategies)
	logging.Log.Infof("lenth: %d", total)
	return nil
}

func CreateStrategyExe(name string) *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"Name": name,
	}, CreateStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(StrategyParam)

	err = checkNameParam(params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, StrategyModule, false)
	if err != nil {
		return err
	}

	strategy := DcsStrategies{
		UUID:         uuid.NewString(),
		Name:         params.Name,
		TimePoint:    params.TimePoint,
		Interval:     params.Interval,
		StrategyType: params.StrategyType,
		Description:  params.Description,
	}
	if err = strategy.Create(); err != nil {
		return err
	}

	return nil
}

func DeleteStrategyExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID": "4e3f0123-86ef-4e84-a4e4-a673db294aea",
		"Name": "",
	}, DeleteStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(BaseParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, StrategyModule, true)
	if err != nil {
		return err
	}

	// 检查是否还在被某个复制关系所绑定

	var strategy DcsStrategies
	strategy.Name = params.Name
	strategy.UUID = params.UUID
	if err = strategy.Delete(); err != nil {
		return err
	}
	return nil
}

func UpdateStrategyExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"UUID":        "",
		"Name":        "relation2",
		"NewName":     "strategy2",
		"Description": "desc1&&",
	}, UpdateStrategyFun)
	if err != nil {
		return err
	}
	params := anyParams.(StrategyUpdateParam)

	err = checkIDAndNameParam(params.UUID, params.Name)
	if err != nil {
		return err
	}

	err = checkNameParam(params.NewName)
	if err != nil {
		return err
	}

	err = checkExisted(params.Name, params.UUID, StrategyModule, true)
	if err != nil {
		return err
	}

	err = checkExisted(params.NewName, "", StrategyModule, false)
	if err != nil {
		return err
	}

	var strategy DcsStrategies
	strategy.UUID = params.UUID
	strategy.Name = params.Name
	update := map[string]interface{}{
		"description": params.Description,
		"name":        params.NewName,
	}
	err = strategy.Update(update)
	if err != nil {
		return err
	}
	return nil
}

func ListStrategiesExe() *fault.Fault {
	anyParams, err := ParamWrap(map[string]interface{}{
		"PageNumber":  "2",
		"PageSize":    "3",
		"SortBy":      "Name",
		"Order":       "asc",
		"FilterBy":    "Name",
		"FilterValue": "",
	}, ListStrategiesFun)
	if err != nil {
		return err
	}
	params := anyParams.(FilterParam)

	var strategyDb DcsStrategies
	strategies, total, err := strategyDb.List(params.FilterBy, params.FilterValue, params.Order, params.SortBy, params.PageSize, params.PageNumber)
	if err != nil {
		return err
	}

	logging.Log.Infof("%+v", strategies)
	logging.Log.Infof("lenth: %d", total)
	return nil
}

func deleteRelation2DBRollback(params RelationParam) *fault.Fault {
	var relation DcsRelations
	relation.UUID = params.UUID
	if err := relation.Delete(); err != nil {
		return err
	}
	return nil
}
