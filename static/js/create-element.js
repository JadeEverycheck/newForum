
function createElement(e) {
    if (typeof e === "string") {
        return document.createTextNode(e)
    }

    const elem = document.createElement(e.tag)

    for (const key in e.properies) {
        elem[key] = e.properies[key]
    }

    if (e.children !== undefined) {
        for (const child of e.children){
            elem.appendChild(createElement(child))
        }
    }

    return elem
}