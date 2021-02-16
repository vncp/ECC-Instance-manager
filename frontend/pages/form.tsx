import Head from "next/head";
import Form from "react-bootstrap/Form";
import MainNavbar from "../components/MainNavbar";
import Button from "react-bootstrap/Button";

const RequestForm = () => {
  return (
    <>
      <Head>
        <title>Request Linux Instance Form</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainNavbar />
      <div
        style={{
          boxShadow: "2px 2px 5px 4px rgba(0, 0, 0, 0.07)",
          marginLeft: "30vw",
          marginRight: "30vw",
          paddingTop: "2vh",
          paddingBottom: "3vh",
          display: "flex",
          marginTop: "15vh",
          flexDirection: "column",
          alignItems: "center",
          backgroundColor: "#FFFFFF",
        }}
      >
        <h2>Linux Instance Request Form</h2>
        <Form style={{ display: "flex", flexDirection: "column" }}>
          <Form.Group controlId="requestFormName">
            <Form.Label>Name</Form.Label>
            <Form.Control type="name" />
          </Form.Group>
          <Form.Group controlId="requestFormNetid">
            <Form.Label>NetID</Form.Label>
            <Form.Control type="netid" />
          </Form.Group>
          <Form.Group controlId="requestFormEmail">
            <Form.Label>Email</Form.Label>
            <Form.Control type="email" />
            <Form.Text className="text-muted">
              We'll use this to notify you that your request has been processed.
            </Form.Text>
          </Form.Group>
          <Form.Group controlId="requestFormReason">
            <Form.Label>Course</Form.Label>
            <Form.Control type="reason" />
            <Form.Text className="text-muted"></Form.Text>
            <Form.Text className="text-muted">
              Enter your course or reason for request.
            </Form.Text>
          </Form.Group>
          <Button style={{ marginTop: "1vh" }} type="submit">
            Submit
          </Button>{" "}
        </Form>
      </div>
    </>
  );
};

export default RequestForm;
