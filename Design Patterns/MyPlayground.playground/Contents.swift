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
print("xxx")



