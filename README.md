<h1 align="center">
  <br>
  p2pdb-log
  <br>
</h1>

基于默克尔-CRDT实现的只允许追加的日志组件
[Merkle-CRDT](https://research.protocol.ai/blog/2019/a-new-lab-for-resilient-networks-research/PL-TechRep-merkleCRDT-v0.1-Dec30.pdf)


### 内容列表 


- [目标](#目标)
	- [使用场景](#使用场景)
- [架构](#架构)
	- [目录分层](#目录分层)

## 背景
支撑P2PDB，实现多种不同数据存储结构的数据库。 基于merkle+CRDT 实现日志的递增存储

## 部分代码参考来源
[go-ipfs-log](https://github.com/berty/go-ipfs-log)



## Licensing

*p2pdb-log* is licensed under the Apache License, Version 2.0.
See [LICENSE](LICENSE) for the full license text.
