# Design Patterns Implementation / 设计模式实现

[English Version](./README_EN.md) | 中文版

This repository contains examples of 22 classic design patterns implemented in Go.
本仓库包含了 Go 语言实现的 22 种经典设计模式示例。

## Introduction / 简介

Design patterns are typical solutions to commonly occurring problems in software design. They can be used to speed up the development process by providing tested, proven development paradigms.
设计模式是软件设计中常见问题的典型解决方案。它们提供了经过测试和验证的开发范式，可以加快开发过程。

## Patterns Implemented / 已实现的模式

### 1. Creational Patterns / 创建型模式
- **Factory Method / 工厂方法**
  - Creates objects without specifying the exact class / 创建对象时不指定具体类
- **Abstract Factory / 抽象工厂**
  - Creates families of related objects / 创建一系列相关对象
- **Builder / 建造者**
  - Constructs complex objects step by step / 分步骤构建复杂对象
- **Prototype / 原型**
  - Creates objects by cloning an existing object / 通过克隆现有对象创建对象
- **Singleton / 单例**
  - Ensures a class has only one instance / 确保一个类只有一个实例

### 2. Structural Patterns / 结构型模式
- **Adapter / 适配器**
  - Allows incompatible interfaces to work together / 使不兼容的接口能够一起工作
- **Bridge / 桥接**
  - Separates an object's interface from its implementation / 将对象的接口与实现分离
- **Composite / 组合**
  - A tree structure of simple and composite objects / 简单和复合对象的树形结构
- **Decorator / 装饰器**
  - Add responsibilities to objects dynamically / 动态地给对象添加功能
- **Facade / 外观**
  - A single interface to a complex subsystem / 为复杂子系统提供统一接口
- **Flyweight / 享元**
  - A fine-grained instance used for efficient sharing / 共享细粒度实例以提高效率
- **Proxy / 代理**
  - An object representing another object / 一个对象代表另一个对象

### 3. Behavioral Patterns / 行为型模式
- **Chain of Responsibility / 责任链**
  - A way of passing a request between a chain of objects / 在对象链中传递请求
- **Command / 命令**
  - Encapsulate a command request as an object / 将命令请求封装为对象
- **Iterator / 迭代器**
  - Sequentially access the elements of a collection / 顺序访问集合元素
- **Mediator / 中介者**
  - Defines simplified communication between classes / 定义类之间的简化通信
- **Memento / 备忘录**
  - Capture and restore an object's internal state / 捕获和恢复对象的内部状态
- **Observer / 观察者**
  - A way of notifying change to a number of classes / 通知多个类关于状态的变化
- **State / 状态**
  - Alter an object's behavior when its state changes / 对象状态改变时改变其行为
- **Strategy / 策略**
  - Encapsulates an algorithm inside a class / 将算法封装在类中
- **Template Method / 模板方法**
  - Defer the exact steps of an algorithm to a subclass / 将算法的具体步骤延迟到子类
- **Visitor / 访问者**
  - Defines a new operation to a class without change / 在不改变类的情况下定义新操作

## Usage / 使用方法

Each pattern has its own directory containing implementation files and test files.
每个模式都有自己的目录，包含实现文件和测试文件。

To run the tests for a specific pattern:
运行特定模式的测试：

```bash
cd design_mod/<pattern_name>
go test -v
```

## Pattern Categories / 模式分类

1. **Creational Patterns / 创建型模式**
   - Focus on object creation mechanisms / 关注对象创建机制
   - Make system independent of creation, composition and representation / 使系统独立于对象的创建、组合和表示

2. **Structural Patterns / 结构型模式**
   - Deal with object composition / 处理对象组合
   - Help ensure that when parts of a system change, the entire structure doesn't need to change / 确保系统部分改变时，整体结构不需要改变

3. **Behavioral Patterns / 行为型模式**
   - Focus on communication between objects / 关注对象之间的通信
   - How objects interact and distribute responsibility / 对象如何交互以及分配职责

## Directory Structure / 目录结构

```
design_mod/
├── abstract_factory/    // 抽象工厂模式
├── adapter/            // 适配器模式
├── bridge/             // 桥接模式
├── builder/            // 建造者模式
├── chain_of_responsibility/  // 责任链模式
├── command/            // 命令模式
├── composite/          // 组合模式
├── decorator/          // 装饰器模式
├── facade/             // 外观模式
├── factory/            // 工厂方法模式
├── flyweight/          // 享元模式
├── iterator/           // 迭代器模式
├── mediator/           // 中介者模式
├── memento/            // 备忘录模式
├── observer/           // 观察者模式
├── prototype/          // 原型模式
├── proxy/              // 代理模式
├── single/             // 单例模式
├── state/              // 状态模式
├── strategy/           // 策略模式
├── template/           // 模板方法模式
└── visitor/            // 访问者模式
```

## Contributing / 贡献

Feel free to contribute more patterns or improve existing implementations.
欢迎贡献更多模式或改进现有实现。

## License / 许可

MIT License
MIT 许可证 