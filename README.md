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
P2PDB数据库在p2pdb-log之上实现，p2pdb-log是一种用于分布式系统的不可变的、基于操作的无冲突复制数据结构 (CRDT)与Merkle DAG（有向无环图）实现。如果所有 P2PDB 数据库类型都不符合您的需求和/或您需要特定于案例的功能，您可以轻松使用日志模块实现您想要的数据库

## 部分代码参考来源
[go-ipfs-log](https://github.com/berty/go-ipfs-log)



## Licensing

*p2pdb-log* is licensed under the Apache License, Version 2.0.
See [LICENSE](LICENSE) for the full license text.
