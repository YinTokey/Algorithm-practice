//: Playground - noun: a place where people can play

import UIKit

/*
 =====================================================
 创建型
 */


/*
 对象模板模式 (初始化时完成赋值)
 */
class Student{
    var name:String?
    var age:Int?
    
    init(name:String,age:Int){
        self.name = name
        self.age = age
    }
}

/*
 原型模式，克隆已有的对象，来创建新的对象，已有的对象即为原型. 比如系统自带的copy方法（要实现NSCoping协议）
 */
protocol Prototype{
    func copy()->Prototype
}

struct Product:Prototype{
    var title:String?
    func copy() -> Prototype {
        return Product(title: title)
    }
}

/*
 单例模式，封装共享资源，提供全局访问点
 */

/*
 对象池模式 ，有些资料可能提及这个模式。是单例的扩展，管理一组可重用对象，比如tableview的cell复用机制
 */

/*
 工厂方法模式：定义创建对象的接口，让子类自己觉得要实例化哪一个工厂类，工厂模式使得创建过程延迟到子类进行
 */


/*
 抽象工厂模式：提供一个接口，能够创建一些列相关或相互依赖的对象。比如NSNumber的一些初始化方法 [NSNumber numberWithBool:YES],[NSNumber numberWithInt:1]
 抽象工厂里包含了很多工厂模式
 */


/*
 建造者模式：将对象的创建和配置分离，引入建造者类来配置对象.
 */
class StartBuck{
    
    init(milk:Bool,sugger:Bool){
        
    }
}

class CoffeBuilder{
    private var sugger:Bool  = false
    private var milk:Bool = false
    
    func setSugger(choice:Bool){
        self.sugger = choice
    }
    
    func setMilk(choice:Bool){
        self.milk = choice
    }
    
    func buildObject(name:String)->StartBuck{
        return StartBuck(milk: milk, sugger: sugger)
        
    }
}

let builder = CoffeBuilder()
builder.setMilk(choice: true)
builder.setSugger(choice: false)

let coffee = builder.buildObject(name: "10A")



/*
 =====================================================
 结构型
 */

/*
 适配器模式Adaper，将一个类不兼容的接口，转换为另一个类兼容的接口，常用于旧代码的复用. 连接结构如  A类 --- 适配器 --- B类
 比如下面例子   Target -- Adapter -- Adaptee
 */
protocol Target {
    var value: String { get }
}

struct Adapter: Target {
    let adaptee: Adaptee
    var value: String {
        return "\(adaptee.value)"
    }
    init(_ adaptee: Adaptee) {
        self.adaptee = adaptee
    }
}

struct Adaptee {
    var value: Int
}


/*
 桥接模式
 */


/*
 装饰者模式
 动态给对象添加额外的职责,比子类更加灵活
 */

protocol Component {
    
    var cost:Int { get }
    
}

protocol Decorator:Component {
    
    var component:Component { get }
    init(_ component:Component)
}

struct Coffee:Component {
    
    var cost:Int
    
}

struct Sugar:Decorator{
    var component: Component
    
    init(_ component: Component) {
        self.component = component
    }
    
    var cost: Int {
        return component.cost + 1
    }

}

struct Milk:Decorator {
    var component: Component
    
    init(_ component: Component) {
        self.component = component
    }
    
    var cost: Int {
        return component.cost + 2
    }
    
}

Milk(Sugar(Coffee(cost:19))).cost

/*
 享元模式
 运用共享技术有效支持大量细粒度对象,就是一套缓存系统. SDWebImage中的缓存设计，就有涉及享元模式
 */

/*
 代理模式
 这个没啥好说的了。。
 */



/*
 行为型
 =====================================================
 
 */


/*
 责任链模式
 避免请求发送者和接受者耦合在一起，让多个对象都可以接受请求，这些对象连成一条链，请求沿着这条链传递，直到有对象处理它为止。最典型的就是iOS的响应链
 */



/*
 命令模式
 将请求封装成对象。最典型的就是 RACCommand,  RxSwift的 Action
 */
protocol Command{
    var operation:()->() {get}
    var backup:String {get}
    func undo()
    
}

struct ConcreteCommand:Command{
    var backup: String
    var operation: () -> ()
    func undo() {
        print(backup)
    }
}

/*
 迭代器模式
 提供一种方法顺序访问一个集合内的元素，但是又不暴露具体实现，比如数组字典的 Block 迭代模式
 */
protocol AbstractIterator {
    func hasNext() -> Bool
    func next() -> Int
}

class ConcreteIterator: AbstractIterator {
    private var currentIndex = 0
    var elements:[Int] = []
    
    func hasNext() -> Bool {
        
    }
    
    func next() -> Int {
        
    }
}

protocol AbstractCollection {
    func makeIterator() -> AbstractIterator
}

class ConcreteCollection:AbstractCollection{
    let iterator = ConcreteIterator()
    func makeIterator() -> AbstractIterator {
        return iterator
    }
    
    func add(_ e:Int){
        iterator.elements.append(e)
    }
}

let c = ConcreteCollection()

let iterator = c.makeIterator()
while iterator.hasNext(){
    print(iterator.next())
}

/*
 中介者模式
 用一个中介对象封装一些列对象之间的交互，中介者使得各个对象之间不相互直接引用，从而起到解耦的作用，比如组件化方案中的Mediator
 */


/*
 备忘录模式
 在不破坏封装性的前提下，捕获一个对象的内部状态，并在对象之外保存这个状态。比如 iOS 中的归档解档
 */


/*
 观察者模式
 定义对象间一对多的关系,当一个对象发送改变时，依赖它的对象会得到通知与更新。比如KVO,RxSwift
 */


/*
 状态模式
 允许对象在内部状态发生改变时改变它的行为，对象看起来好像修改了它的类。
 更通俗的解释就是 当一个对象的状态不一样的时候，它调用同一个方法会产生不同的结果。
 */
protocol State{
    func operation()
}

class ConcreteStateA:State{
    func operation() {
        print("A")
    }
}

class ConcreteStateB:State{
    func operation() {
        print("B")
    }
}

class Context{
    var state:State = ConcreteStateA()
    func doSomething(){
        state.operation()
    }
}

let ct = Context()
ct.doSomething()
ct.state = ConcreteStateB()
ct.doSomething()

/*
 策略模式
 就行不同的地图交通路线，可以相互替换，最终都能到达目的地。
 定义一系列算法，把他们都封装起来，并且可以相互替换。
 */


/*
 模板方法模式
 在抽象类中声明方法，在子类中具体实现。比如NSOperation是抽象类，使用的时候一般使用它的子类
 */
