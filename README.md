# red-black-trees

### 左旋操作
节点向左移动,右节点上升到当前节点位置，右节点的左节点变成当前节点的右节点。
1. A的父节点的子节点指针指向C（需要区分左右）
2. A的右节点指向D
3. C的左节点指向A
![](https://tva1.sinaimg.cn/large/008i3skNgy1gqzq7hq5zzj30ym0l813f.jpg)


### 右旋操作
节点向右移动,左节点上升到当前节点位置，左节点的右节点变成当前节点的左节点。
1. A的父节点的子节点指向指向B(需要区分左右)
2. A节点左指针指向E
3. B右节点指向E
![](https://tva1.sinaimg.cn/large/008i3skNgy1gqzq8d5o8hj315m0m814a.jpg)


## 整体流程
### 插入操作
![](https://tva1.sinaimg.cn/large/008i3skNgy1gr0ul09vtaj30u011z14j.jpg)

参考引用：
https://segmentfault.com/a/1190000022278733
https://www.huximi.com/2020/05/13/red-black-tree/
