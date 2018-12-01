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



