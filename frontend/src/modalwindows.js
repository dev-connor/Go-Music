import React from "react";
import { Modal, ModalHeader, ModalBody } from "reactstrap";

export class BuyModalWindow extends React.Component {
    render() {
        return (
            
            <Modal id="buy" tabIndex="-1" role="dialog" isOpen={this.props.showModal} toggle={this.props.toggle}>
                <div role="document">
                    <ModalHeader toggle={this.props.toggle} className="bg-success text-white">
                        Buy Item
                    </ModalHeader>
                    <ModalBody>
                        {/* 신용카드 결제 폼 */}
                    </ModalBody>
                </div>
            </Modal>
        )
    }
}