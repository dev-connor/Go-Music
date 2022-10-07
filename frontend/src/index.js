import React from "react"
import ReactDOM from "react-dom"

class Card extends React.Component {
  render() {
    return (

      <div className="col-md-6 col-lg-4 d-flex align-items-stretch">
        <div className="card mb-3">
          <img className="card-img-top" src={this.props.img} alt={this.props.imgalt}/>
          <div className="card-body">
            <h4 className="card-title">{this.props.id}. {this.props.productName}</h4>
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
  constructor(props) {
    // 부모 컴포넌트로 props 전달
    super(props) 
    // 컴포넌트의 state 객체 초기화
    this.state = {
      cards: []
    }
  }

  componentDidMount() {
    fetch('cards.json')
    .then(res => res.json())
    .then(result => {
      this.setState({
        cards: result
      })
    })  
  }

  render() {
    const cards = this.state.cards
    let items = cards.map(card => <Card {...card}/>)
    
    return(
      <div className="container pt-4">
        <h3 className="'text-center text-primary">Products</h3>
        <div className="pt-4 row">{items}</div>
      </div>
    )
  }
}

ReactDOM.render(<CardContainer/>, document.getElementById('root'))
