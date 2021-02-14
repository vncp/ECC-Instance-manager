import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";

const MainNavbar = () => {
  return (
    <Navbar collapseOnSelect expand="xl" bg="dark" variant="dark" fixed="top">
      <Navbar.Brand href="#home">Remote Linux Instance Management</Navbar.Brand>
      <Navbar.Toggle aria-controls="responsive-navbar-nav" />
      <Navbar.Collapse id="Responsive-navbar-nav">
        <Nav className="ml-auto" fill>
          <Nav.Item>
            <Nav.Link href="/login">Login</Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link href="https:/remote.engr.unr.edu/">Remote Home</Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link href="#Help">Documentation</Nav.Link>
          </Nav.Item>
          <Nav.Item>
            <Nav.Link href="/management">Management</Nav.Link>
          </Nav.Item>
        </Nav>
      </Navbar.Collapse>
    </Navbar>
  );
};

export default MainNavbar;
