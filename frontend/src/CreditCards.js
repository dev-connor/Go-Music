import React from "react";
import {injectStripe, StripeProvider, Elements, CardElement} from 'react-stripe-elements'

const INITIALSTATE = "INITIAL", SUCCESSSTATE = "COMPLETE", FAILEDSTATE = "FAILED"

class CreditCardForm extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            status: INITIALSTATE
        }
    }
    renderCreditCardInformation() {}
    renderSuccess() {}
    renderFailure() {}

    render() {
        let body = null
        switch (this.state.status) {
            case SUCCESSSTATE:
                body = this.renderSuccess()
                break
            case FAILEDSTATE:
                body = this.renderFailure()
                break
            default:
                body = this.renderCreditCardInformation()
        }
        return (
            <div>
                {body}
            </div>
        )
    }
    
}