#let project(title: "", authors: (), logo: none, body) = {
  set document(author: authors, title: title)
  set page(numbering: "1", number-align: center)

  let body-font = "CMU Serif"
  let sans-font = "CMU Sans Serif"

  set text(font: body-font, 12pt, lang: "ru")
  show heading: set text(font: sans-font)
  set heading(numbering: "1.1.")

  v(0.6fr)
  if logo != none {
    align(right, image(logo, width: 26%))
  }
  v(9.6fr)

  text(font: sans-font, 2em, weight: 700, title)

  pad(
    top: 0.7em,
    right: 20%,
    grid(
      columns: (1fr,) * calc.min(3, authors.len()),
      gutter: 1em,
      ..authors.map(author => align(start, strong(author))),
    ),
  )

  v(2.4fr)
  pagebreak()

  outline(depth: 2)
  pagebreak()

  set par(justify: true)

  body
}


#show: project.with(
  title: "Отчет по  лабораторным работам по предмету Ибизи",
  authors: (
    "Ковешнииков Даниил 451 группа",
  ),
)

#include "task2/content.typ"

#include "task3/contetnt.typ"

#include "task5/content.typ"

#include "task6/content.typ"
