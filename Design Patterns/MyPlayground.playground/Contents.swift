//: Playground - noun: a place where people can play

import UIKit

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



