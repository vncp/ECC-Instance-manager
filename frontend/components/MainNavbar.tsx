import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Router from "next/router";
import { Logout } from "../services/login_service";
import {
  redirectToLogin,
  redirectToRequestForm,
} from "../services/redirect_service";

const MainNavbar = (props) => {
  return (
    <Navbar collapseOnSelect expand="xl" bg="dark" variant="dark" fixed="top">
      <Navbar.Brand href="/">Remote Linux Instance Management</Navbar.Brand>
      <Navbar.Toggle aria-controls="responsive-navbar-nav" />
      <Navbar.Collapse id="Responsive-navbar-nav">
        <Nav className="ml-auto" fill>
          {!props.loggedIn ? (
            <Nav.Item>
              <Nav.Link onClick={() => redirectToLogin()}>Login</Nav.Link>
            </Nav.Item>
          ) : (
            <Nav.Item>
              <Nav.Link onClick={() => Logout()}>Logout</Nav.Link>
            </Nav.Item>
          )}
          {props.staff ? (
            <Nav.Item>
              <Nav.Link onClick={() => Router.push("/requests")}>
                Requests
              </Nav.Link>
            </Nav.Item>
          ) : (
            <></>
          )}
          <Nav.Item>
            <Nav.Link onClick={() => Router.push("/dashboard")}>
              Manage
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link href="https:/remote.engr.unr.edu/" target="_blank">
              Remote Home
            </Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link
              href="https://ph.engr.unr.edu/w/remote-instance-help/"
              target="_blank"
            >
              Help
            </Nav.Link>
          </Nav.Item>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
  );
};

export default MainNavbar;
