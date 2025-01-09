# Design Patterns Implementation

[中文版](./README.md) | English Version

This repository contains examples of 22 classic design patterns implemented in Go.

## Introduction

Design patterns are typical solutions to commonly occurring problems in software design. They can be used to speed up the development process by providing tested, proven development paradigms.

## Patterns Implemented

### 1. Creational Patterns
- **Factory Method**
  - Creates objects without specifying the exact class
- **Abstract Factory**
  - Creates families of related objects
- **Builder**
  - Constructs complex objects step by step
- **Prototype**
  - Creates objects by cloning an existing object
- **Singleton**
  - Ensures a class has only one instance

### 2. Structural Patterns
- **Adapter**
  - Allows incompatible interfaces to work together
- **Bridge**
  - Separates an object's interface from its implementation
- **Composite**
  - A tree structure of simple and composite objects
- **Decorator**
  - Add responsibilities to objects dynamically
- **Facade**
  - A single interface to a complex subsystem
- **Flyweight**
  - A fine-grained instance used for efficient sharing
- **Proxy**
  - An object representing another object

### 3. Behavioral Patterns
- **Chain of Responsibility**
  - A way of passing a request between a chain of objects
- **Command**
  - Encapsulate a command request as an object
- **Iterator**
  - Sequentially access the elements of a collection
- **Mediator**
  - Defines simplified communication between classes
- **Memento**
  - Capture and restore an object's internal state
- **Observer**
  - A way of notifying change to a number of classes
- **State**
  - Alter an object's behavior when its state changes
- **Strategy**
  - Encapsulates an algorithm inside a class
- **Template Method**
  - Defer the exact steps of an algorithm to a subclass
- **Visitor**
  - Defines a new operation to a class without change

## Usage

Each pattern has its own directory containing implementation files and test files.

To run the tests for a specific pattern:

```bash
cd design_mod/<pattern_name>
go test -v
```

## Pattern Categories

1. **Creational Patterns**
   - Focus on object creation mechanisms
   - Make system independent of creation, composition and representation

2. **Structural Patterns**
   - Deal with object composition
   - Help ensure that when parts of a system change, the entire structure doesn't need to change

3. **Behavioral Patterns**
   - Focus on communication between objects
   - How objects interact and distribute responsibility

## Directory Structure

```
design_mod/
├── abstract_factory/    // Abstract Factory Pattern
├── adapter/            // Adapter Pattern
├── bridge/             // Bridge Pattern
├── builder/            // Builder Pattern
├── chain_of_responsibility/  // Chain of Responsibility Pattern
├── command/            // Command Pattern
├── composite/          // Composite Pattern
├── decorator/          // Decorator Pattern
├── facade/             // Facade Pattern
├── factory/            // Factory Method Pattern
├── flyweight/          // Flyweight Pattern
├── iterator/           // Iterator Pattern
├── mediator/           // Mediator Pattern
├── memento/            // Memento Pattern
├── observer/           // Observer Pattern
├── prototype/          // Prototype Pattern
├── proxy/              // Proxy Pattern
├── single/             // Singleton Pattern
├── state/              // State Pattern
├── strategy/           // Strategy Pattern
├── template/           // Template Method Pattern
└── visitor/            // Visitor Pattern
```

## Contributing

Feel free to contribute more patterns or improve existing implementations.

## License

MIT License 