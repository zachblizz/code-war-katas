function lookAndSaySequence(firstElement, n) {
  return helper(firstElement, 1, n)
}

function helper(elements, n, stop) {
  if (n >= stop || elements === "" || stop === 1) {
    return elements
  }

  let count = 1
  let currElem = ""
  const newElements = []
  const elems = elements.split("")
  
  elems.forEach((e, i, arr) => {
    const atEnd = i === arr.length - 1

    if (currElem === "") {
      currElem = e
    }

    if (i !== 0 && currElem === e) {
      count++

      if (atEnd) {
        newElements.push(`${count}${currElem}`)
      }
    } else if ((currElem !== e || atEnd) && stop > 1) {
      newElements.push(`${count}${currElem}`)
      count = 1
      currElem = e

      if (atEnd && arr.length > 1) {
        newElements.push(`${count}${currElem}`)
      }
    }
  })

  return helper(newElements.join(""), ++n, stop)
}

module.exports = lookAndSaySequence
