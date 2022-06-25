package internal

import "awesomeProject/pkg/HOM"

var AppElements = HOM.NewHOMElement(
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
)
