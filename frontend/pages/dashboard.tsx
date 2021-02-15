import React, { useEffect } from "react";
import MainNavbar from "../components/MainNavbar";
import { AuthProps, privateRoute } from "../components/privateRoute";
import { AuthToken } from "../services/auth_token";
import getInstances from "../services/instances_service";

type Props = AuthProps;

const Dashboard = ({ auth }: Props) => {
  useEffect(() => getInstances(auth.authorizationString));
  return (
    <>
      <MainNavbar />
      <div style={{ marginTop: "3vh" }}>
        <h1>Management Page</h1>
        <p>
          <strong>user</strong>: {auth.decodedToken.netid}
        </p>
        <p>
          <strong>isValid</strong>: {auth.isValid.toString()}
        </p>
        <p>
          <strong>isExpired</strong>: {auth.isExpired.toString()}
        </p>
        <p>
          <strong>authorizationString</strong>: {auth.authorizationString}
        </p>
        <p>
          <strong>expiresAt</strong>: {auth.expiresAt.toString()}
        </p>
      </div>
    </>
  );
};

export default privateRoute(Dashboard);
