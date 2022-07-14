import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Form, Container, Modal } from 'react-bootstrap'
import Transport from './single-transport.component';

const Transports = () => {

    const [transports, setTransports] = useState([])
    const [refreshData, setRefreshData] = useState(false)

    const [changeTransport, setChangeTransport] = useState({"change": false, "id": 0})

    const [addNewTransport, setAddNewTransport] = useState(false)
    const [newTransport, setNewTransport] = useState({"name": "", "description": "", "modality": ""})

    //gets run at initial loadup
    useEffect(() => {
        getAllTransports();
    }, [])

    //refreshes the page
    if(refreshData){
        setRefreshData(false);
        getAllTransports();
    }

    return (
        <div>
            
        {/* add new transport button */}
        <Container>
            <Button onClick={() => setAddNewTransport(true)}>Add new transport</Button>
        </Container>

        {/* list all transports */}
        <Container>
            {transports != null && transports.map((transport, i) => (
                <Transport transportData={transport} deleteTransport={deleteSingleTransport} setChangeTransport={setChangeTransport}/>
            ))}
        </Container>
        
        {/* popup for adding a new transport */}
        <Modal show={addNewTransport} onHide={() => setAddNewTransport(false)} centered>
            <Modal.Header closeButton>
                <Modal.Title>Add Order</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form.Group>
                    <Form.Label >name</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.name = event.target.value}}/>
                    <Form.Label>description</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.description = event.target.value}}/>
                    <Form.Label >modality</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.modality = event.target.value}}/>
                </Form.Group>
                <Button onClick={() => addSingleTransport()}>Add</Button>
                <Button onClick={() => setAddNewTransport(false)}>Cancel</Button>
            </Modal.Body>
        </Modal>
        
        {/* popup for changing a transport */}
        <Modal show={changeTransport.change} onHide={() => setChangeTransport({"change": false, "id": 0})} centered>
            <Modal.Header closeButton>
                <Modal.Title>Change Transport</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form.Group>
                    <Form.Label >name</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.name = event.target.value}}/>
                    <Form.Label>description</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.description = event.target.value}}/>
                    <Form.Label >modality</Form.Label>
                    <Form.Control onChange={(event) => {newTransport.modality = event.target.value}}/>
                </Form.Group>
                <Button onClick={() => changeSingleTransport()}>Change</Button>
                <Button onClick={() => setChangeTransport({"change": false, "id": 0})}>Cancel</Button>
            </Modal.Body>
        </Modal>
    </div>
    )

    //changes the transport
    function changeSingleTransport(){
        changeTransport.change = false;
        var url = "http://localhost:8080/transport/" + changeTransport.id
        axios.put(url, newTransport)
            .then(response => {
            if(response.status == 200){
                setRefreshData(true)
            }
        })
    }
    
    //creates a new order
    function addSingleTransport(){
        setAddNewTransport(false)
        var url = "http://localhost:8080/transport"
        axios.post(url, {
            "name": newTransport.name,
            "description": newTransport.description,
            "modality": newTransport.modality,
        }).then(response => {
            if(response.status == 200){
                setRefreshData(true)
            }
        })
    }
    
    //gets all the transports
    function getAllTransports(){
        var url = "http://localhost:8080/transports"
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if(response.status == 200){
                console.log(response.data)
                setTransports(response.data)
            }
        })
    }
    
    //deletes a transport
    function deleteSingleTransport(id){
        var url = "http://localhost:8080/transport/" + id
        axios.delete(url, {

        }).then(response => {
            if(response.status == 200){
                setRefreshData(true)
            }
        })
    }
}

export default Transports