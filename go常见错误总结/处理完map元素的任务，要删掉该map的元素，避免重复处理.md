> map 处理完map元素的任务，要删掉该map的元素，避免重复处理

var nodeMap map[string]*MinLimit = make(map[string]*MinLimit, 10)

for {
		
        // 如果 nodeMap的某个节点信息已经被处理，node需要从nodeMap上删除
        
        // 从nodeMap中删除该节点
        delete(nodeMap, node)
		
}