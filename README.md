# Go 设计模式

1. 设计模式使人们可以更加简单方便地复用成功的设计和体系结构，将已经证实的技术表述成设计模式也会使新系统开发者更加容易理解其设计思路，是形式化的最佳实践
2. 设计模式能够帮助做出有利于系统复用的选择，避免设计损害了系统的复用性
3. 设计模式的核心在于为相关问题提供解决方案(提供重复出现的设计问题的解决方案，以编写灵活和可复用的面向对象软件
4. 设计模式代表了有经验的面向对象软件开发人员所使用的最佳实践，设计模式是软件开发人员在软件开发过程中所面临的一般问题的解决方案，这些解决方案是由许多软件开发人员经过一段实践的试验和探索获得的
5. 软件设计模式是在软件设计中给定上下文中对常见问题的通用可重用解决方案，它不是可以直接转换为源代码或机器代码的最终设计，它是如何解决在不同情况下使用的问题的描述模板，设计模式是形式化的最佳实践，程序员可以可以在设计应用程序或系统时用来解决常见问题
6. 一般而言，一个设计模式有四个基本要素
   - 模式名称
   - 问题：描述在何时使用模式
   - 解决方案：描述了设计的组成成分，它们之间的相关关系及各自的职责和协作方式
   - 效果：藐视了模式应用的效果及使用模式应权衡的问题
7. 设计模式根据其目的可分为创建型、结构型、行为型三种
   - 创建型：与对象的创建有关
   - 结构型：处理类或对象的组合
   - 对类和对象怎样交互和怎样分配职责进行描述
8. 设计模式根据其作用范围可以分为：类型、对象型两种
   - 类型：处理类和子类的关系，这些关系通过继承建立，是静态的
   - 对象型：处理对象间的关系，这些关系在运行时刻是可以变化的，是动态的
9. 对比
   - 创建型类模式将对象的部分创建工作延迟到子类，创建型对象模式则将对象的部分创建工作延迟到另一个对象中
   - 结构型类模式使用继承机制来组合类，而结构型对象模式则描述了对象的组装方式
   - 行为型模式使用继承描述算法和控制流，而行为型对象模式则描述一组对象怎样协作完成单个对象所无法完成的任务

## 创建型
  - [单例模式](./01-Abstract_Factory/abstract_factory.go)
  - [工厂模式](./02-Builder/builder.go)
  - [建造者模式](./03-Factory_Method/factory_method.go)
  - [原型模式](./04-Prototype/prototype.go)
  - [单例模式](./05-Singleton/singleton.go)
## 结构型
  - [适配器模式](./06-Adapter/adapter.go)
  - [桥接模式](./07-Bridge/bridge.go)
  - [组合模式](./08-Composite/composite.go)
  - [装饰者模式](./09-Decorator/decorator.go)
  - [外观模式](./10-Facade/facade.go)
  - [享元模式](./11-Flyweight/flyweight.go)
  - [代理模式](./12-Proxy/proxy.go)
## 行为型
  - [职责链模式](./13-Chain_of_Responsibility/chain_of_responsibility.go)
  - [命令模式](./14-Command/command.go)
  - [解释器模式](./15-Interpreter/interpreter.go)
  - [迭代器模式](./16-Iterator/iterator.go)
  - [中介模式](./17-Mediator/mediator.go)
  - [备忘录模式](./18-Memento/memento.go)
  - [观察者模式](./19-Observer/observer.go)
  - [状态模式](./20-State/state.go)
  - [策略模式](./21-Strategy/strategy.go)
  - [模板模式](./22-Template_method/template_method.go)
  - [访问者模式](./23-Visitor/visitor.go)
