import React from "react"
import ReactDOM from "react-dom"

class Card extends React.Component {
  render() {
    return (

      <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
        <div className="card mb-3">
          <img className="card-img-top" src={this.props.img} alt={this.props.imgalt}/>
          <div className="card-body">
            <h4 className="card-title">{this.props.productName}</h4>
              Price: <strong>{this.props.price}</strong>
            <p className="card-text">{this.props.desc}</p>
            <a href="#" className="btn btn-primary">Buy</a>
          </div>
        </div>
      </div>
    )
  }
}

class CardContainer extends React.Component {
  render() {
    return(
      
      <div>
        <Card key='1' img="img/strings.png" alt="strings" productName="Strings" price='100.0' desc="A very authentic and beautiful instrument!!"/>
        <Card key='2' img="img/redguitar.jpeg" alt="redg" productName="Red Guitar" price='299.0' desc="A really cool red guitar that can produce super cool music!!"/>
      </div>
    )
  }
}

ReactDOM.render(<CardContainer/>, document.getElementById('root'))
