package helpers

func ArrTree(arr []string,pid int ,pidField string,child string ,addChild bool )  {

//	       if ($this->isNewRecord) {
	//            if ($this->pid == 0) {
	//                $this->tree = TreeHelper::defaultTreeKey();
	//            } else {
	//                list($level, $tree) = $this->getParentData();
	//                $this->level = $level;
	//                $this->tree = $tree;
	//            }
	//        } else {
	//            // 修改父类
	//            if (isset($this->oldAttributes['pid']) && $this->oldAttributes['pid'] != $this->pid) {
	//                list($level, $tree) = $this->getParentData();
	//                // 查找所有子级
	//                $list = self::find()
	//                            ->where(['like', 'tree', $this->tree . TreeHelper::prefixTreeKey($this->id) . '%', false])
	//                            ->select(['id', 'level', 'tree', 'pid'])
	//                            ->asArray()
	//                            ->all();
	//
	//                $distanceLevel = $level - $this->level;
	//                // 递归修改
	//                $data = ArrayHelper::itemsMerge($list, $this->id);
	//                $this->recursionUpdate($data, $distanceLevel, $tree);
	//
	//                $this->level = $level;
	//                $this->tree = $tree;
	//            }
	//        }
}