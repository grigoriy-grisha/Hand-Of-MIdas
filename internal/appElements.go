package internal

import "awesomeProject/pkg/HOM"

//
//import "awesomeProject/pkg/HOM"
//

var AppElements = HOM.NewHOMElement(
	HOM.NewElementParams{
		Style: &HOM.Style{
			PaddingRight:    5,
			PaddingLeft:     5,
			PaddingTop:      2,
			PaddingBottom:   2,
			VerticalContent: HOM.VerticalContentCenter,
			AlignContent:    HOM.AlignContentCenter,
			AlignItems:      HOM.AlignItemsEnd,
			Border:          true,
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
							Width:         "50%",
						},
						Children: &HOM.Children{
							Elements: []*HOM.Element{
								HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
								HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
								HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
							},
						},
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
																		HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
																		HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
																		HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
																		HOM.NewHOMElement(HOM.NewElementParams{Text: &HOM.Text{Value: "hello"}}),
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
										Text: &HOM.Text{Value: "hello world "},
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
