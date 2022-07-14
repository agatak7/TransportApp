import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import {Button, Card, Row, Col} from 'react-bootstrap'

const Transport = ({transportData, deleteTransport, setChangeTransport}) => {
    return (
        <Card>
            <Row>
                <Col>Name:{ transportData !== undefined && transportData.name}</Col>
                <Col>Description:{ transportData !== undefined && transportData.description}</Col>
                <Col>Modality:{ transportData !== undefined && transportData.modality}</Col>
                <Col><Button onClick={() => deleteTransport(transportData._id)}>delete transport</Button></Col>
                <Col><Button onClick={() => changeTransport()}>edit transport</Button></Col>
            </Row>
        </Card>
    )

    function changeTransport(){
        setChangeTransport({
            "change": true,
            "id": transportData._id
        })
    }
}
export default Transport