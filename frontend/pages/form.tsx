import Head from "next/head";
import Form from "react-bootstrap/Form";
import MainNavbar from "../components/MainNavbar";
import styles from "../styles/RequestForm.module.css";

const RequestForm = () => {
  return (
    <>
      <Head>
        <title>Request Linux Instance Form</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainNavbar />
      <div className={styles.container}>
        <h1>Remote Linux Instance Request Form</h1>
        <h1 />
        <Form>
          <Form.Group controlId="requestFormName">
            <Form.Label>Name</Form.Label>
            <Form.Control type="name" placeholder="Name" />
          </Form.Group>
          <Form.Group controlId="requestFormNetid">
            <Form.Label>Netid</Form.Label>
            <Form.Control type="netid" placeholder="Netid" />
          </Form.Group>
          <Form.Group controlId="requestFormEmail">
            <Form.Label>Email</Form.Label>
            <Form.Control type="email" placeholder="example@nevada.unr.edu" />
            <Form.Text className="text-muted">
              We'll use this to notify you that your request has been processed.
            </Form.Text>
          </Form.Group>
          <Form.Group controlId="requestFormReason">
            <Form.Label>Course</Form.Label>
            <Form.Control
              type="reason"
              placeholder="Course or reason for request"
            />
            <Form.Text className="text-muted"></Form.Text>
          </Form.Group>
        </Form>
      </div>
    </>
  );
};

export default RequestForm;
