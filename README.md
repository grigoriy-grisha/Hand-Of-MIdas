# Простейшая реализация коносльного рендеринга из древовидной структуры, подобной HTML

![1052669](https://user-images.githubusercontent.com/67334706/175805763-e1545c4c-37af-4b0e-80da-b691f0726e61.jpg)


![Screenshot_58](https://user-images.githubusercontent.com/67334706/175764324-d611c2eb-0110-4fc5-a9be-1d41d082dc03.png)

Проект состоит из 3 модулей

- HOM (Hand OF Midas) - модуль занимется высчитыванием границ элементов, размеров. Модуль разбит на две фазы. 1-высчитывание размеров элементов, 2- высчитывание границ элементова основываясь на их размерах
- HOMR (Hand OF Midas Renderer) - Модуль занмиется отображением всего на экран, имея размеры и границы элементов, так же модуль ренедрит текст во всех 9 выриациях
- HOMF (Hand OF Midas Framework) - Простая оболчка над HOM, Представляет из себя фасад для взаимодействия с бибилиотекой, позволяет отлавливать клики на элементах и запускать перерендер всего содержимого 

Проект очень простой и сырой, разрабатывался с целью развлечения и получения навыков, возможно, когда-то будет доведен до ума, и сможет рендерить HTML/CSS в консоли

Так же, хоть все и работает шустро, тут можно выделить подобие браузерному рендерингу. Все разделено на фазы.
Можно делать минимум работы для изменения цвета, либо половину работы при изменении позиционирования только одного элемента (Вынос из основного потока)

Построенное дерево выглядит так:

```go
HOM.NewHOMElement(
  HOM.NewElementParams{
    Style: &HOM.Style{
      PaddingRight:     5,
      PaddingLeft:      5,
      PaddingTop:       2,
      PaddingBottom:    2,
      VerticalContent:  HOM.VerticalContentCenter,
      AlignContent:     HOM.AlignContentCenter,
      AlignItems:       HOM.AlignItemsEnd,
      ContentDirection: HOM.VerticalDirection,
      Border:           true,
    },
    Children: &HOM.Children{
      Elements: []*HOM.Element{
        HOM.NewHOMElement(
          HOM.NewElementParams{
            Style: &HOM.Style{
              PaddingRight:    5,
              PaddingLeft:     5,
              PaddingTop:      2,
              PaddingBottom:   2,
              Border:          true,
              Width:           "100%",
              AlignContent:    HOM.AlignContentCenter,
              VerticalContent: HOM.VerticalContentCenter,
            },
            Text: &HOM.Text{Value: "Hello from the most simple and crippled console renderer in the world!"},
          },
        ),
        HOM.NewHOMElement(
          HOM.NewElementParams{
            Style: &HOM.Style{
              PaddingRight:     5,
              PaddingLeft:      5,
              PaddingTop:       2,
              PaddingBottom:    2,
              Border:           true,
              ContentDirection: HOM.VerticalDirection,
            },
            Children: &HOM.Children{
              Elements: []*HOM.Element{
                HOM.NewHOMElement(
                  HOM.NewElementParams{
                    Style: &HOM.Style{
                      PaddingRight:  5,
                      PaddingLeft:   5,
                      PaddingTop:    2,
                      PaddingBottom: 2,
                      Border:        true,
                    },
                    Children: &HOM.Children{
                      Elements: []*HOM.Element{
                        HOM.NewHOMElement(
                          HOM.NewElementParams{
                            Style: &HOM.Style{
                              ContentDirection: HOM.VerticalDirection,
                            },
                            Children: &HOM.Children{Elements: []*HOM.Element{
                              HOM.NewHOMElement(
                                HOM.NewElementParams{
                                  Children: &HOM.Children{Elements: []*HOM.Element{
                                    HOM.NewHOMElement(HOM.NewElementParams{ID: "1", Text: &HOM.Text{Value: "hello"}}),
                                    HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
                                    HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
                                    HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
                                  }},
                                },
                              ),
                              HOM.NewHOMElement(
                                HOM.NewElementParams{
                                  Children: &HOM.Children{Elements: []*HOM.Element{
                                    HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "Some Text"}}),
                                  }},
                                },
                              ),
                            },
                            },
                          },
                        ),
                      },
                    },
                  },
                ),
                HOM.NewHOMElement(
                  HOM.NewElementParams{
                    Style: &HOM.Style{
                      PaddingRight:  5,
                      PaddingLeft:   5,
                      PaddingTop:    2,
                      PaddingBottom: 2,
                    },
                    Text: &HOM.Text{Value: "Нажми на меня и я изменю текст"},
                    OnClick: func(element *HOM.Element) {
                      element.Text.Value = "Текст изменен"
                    },
                  },
                ),
              },
            },
          },
        ),
      },
    },
  },
)```
