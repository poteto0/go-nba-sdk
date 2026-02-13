---
name: go-development
description: Use for go/golang coding.
---

# Golang Coding

## Overview

This skill define go/golang coding.
References latest golang's features w/o web-searching.

## v1.26

- The built-in new function,

  ```go
  type Player struct {
    Name string
    BestScore *int
  }

  func newPlayer(name string, score int) *Player {
    return &Player{
      Name: name,
      BestScore: new(score),
    }
  }
  ```

- The restriction that a generic type

  ```go
  type Adder[A Adder[A]] interface {
    Add(A) A
  }

  func algo[A Adder[A]](x, y A) A {
    return x.Add(y)
  }
  ```
