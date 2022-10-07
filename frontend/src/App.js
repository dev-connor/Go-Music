
import React from "react";
import CardContainer from "./ProductCards";
import Nav from './Navigation'
import {SignInModalWindow, BuyModalWindow} from './modalwindows'
import About from "./About";
// import Orders from './orders'
import {BrowserRouter as Router, Route} from "react-router-dom"

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user: {
        loggedin: false,
        name: "",
      }
    };
    // this.showSignInModalWindow = this.showSignInModalWindow.bind(this);
    // this.toggleSignInModalWindow = this.toggleSignInModalWindow.bind(this);
    // this.showBuyModalWindow = this.showBuyModalWindow.bind(this);
    // this.toggleBuyModalWindow = this.toggleBuyModalWindow.bind(this);
  }

  handleSignedIn(user) {
    this.setState({
      user: user
    });
  }

  showSignInModalWindow() {
    const state = this.state
    const newState = Object.assign({}.state,{showSignInModalWindow:true})
    this.setState(newState)
  }

  showBuyModalWindow(id, price) {
    const state = this.state
    const newState = Object.assign({}, state, {showBuyModal:true, productid: id, price: price})
    this.setState(newState)
  }

  toggleSignInModalWindow() {
    // const start = this.state
    // const newState = Object.assign({}, state, {showSignInModal:!state.showBuyModal})
    // this.setState(newState)
  }
  
  componentDidMount() {
    fetch('user.json')
      .then(res => res.json())
      .then((result) => {
        console.log('Fetch...');
        this.setState({
          user: result
        });
      });
  }

  render() {
  return (
    <div>
      <Router>
        <div>
          <Nav user={this.state.user} showModalWindow={this.showSignInModalWindow}/>
            <div className="container pt-4 mt-4">
              <Route exact path="/" render={() => <CardContainer location='cards.json' showBuyModal={this.showBuyModalWindow}/>}/>
              <Route exact path="/promos" render={() => <CardContainer location='promos.json' promo={true} showBuyModal={this.showBuyModalWindow}/>}/>
              <Route path="/about" Component={About}/>
            </div>
        </div>
      </Router>
    </div>
  )
}
}

export default App;
