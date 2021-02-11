import { useState, useEffect } from "react";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import RequestTable from "../components/Table";
import Spinner from "react-bootstrap/Spinner";

const BACKEND_URL = "127.0.0.1:3001";

const requestsTest = [
  {
    name: "Vincent Pham",
    netid: "vpham",
    email: "vpham@nevada.unr.edu",
    course: "CS999",
    status: "Unresolved",
    date: "2/10/21",
  },
  {
    name: "Andrew McIntyre",
    netid: "amcintyre",
    email: "amcintyre@nevada.unr.edu",
    course: "CS999",
    status: "Resolved",
    date: "8/15/20",
  },
];

export default function Home() {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [requests, setRequests] = useState([]);

  useEffect(() => {
    fetch(BACKEND_URL)
      .then((res: Response) => res.json())
      .then((result) => {
        setIsLoaded(true);
        setRequests(result);
      }),
      (error) => {
        setIsLoaded(true);
        setError(error);
        console.error(error);
      };
  });

  if (error) {
    return <div className={styles.container}>Error: {error.message}</div>;
  } else if (!isLoaded) {
    return (
      <div className={styles.container}>
        <Head>
          <title>Loading...</title>
        </Head>
        <Spinner animation="border"></Spinner>
        <div>Loading...</div>
      </div>
    );
  } else {
    return (
      <div className={styles.container}>
        <Head>
          <title>Remote Instance Management</title>
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <Navbar
          collapseOnSelect
          expand="xl"
          bg="dark"
          variant="dark"
          fixed="top"
        >
          <Navbar.Brand href="#home">
            Remote Linux Instance Management
          </Navbar.Brand>
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />
          <Navbar.Collapse id="Responsive-navbar-nav">
            <Nav className="ml-auto" fill>
              <Nav.Item>
                <Nav.Link href="https:/remote.engr.unr.edu/">
                  Remote Home
                </Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link href="#Help">Documentation</Nav.Link>
              </Nav.Item>
            </Nav>
          </Navbar.Collapse>
        </Navbar>
        <Container fluid>
          <RequestTable requests={requests}></RequestTable>
        </Container>
      </div>
    );
  }
}
