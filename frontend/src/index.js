import React from "react"
import ReactDOM from "react-dom"

class Card extends React.Component {
  render() {
    const img = "img/strings.png"
    const imgalt = "string"
    const desc = "A very authentic and beautiful instrument!!"
    const price = 100
    const productName = "Strings"

    return (
      <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
        <div className="card mb-3">
          <img className="card-img-top" src={img} alt={imgalt}/>
          <div className="card-body">
            <h4 className="card-title">{productName}</h4>
              Price: <strong>{price}</strong>
            <p className="card-text">{desc}</p>
            <a href="#" className="btn btn-primary">Buy</a>
          </div>
        </div>
      </div>
    )
  }
}

ReactDOM.render(<Card/>, document.getElementById('root'))
