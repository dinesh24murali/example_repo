import React from "react";
const LeftNavComponent = React.lazy(() => import("LeftNav/LeftNav"));
const RightNavComponent = React.lazy(() => import("RightNav/RightNav"));

export default function App() {
  return (
    <div>
      <h1>Shell</h1>
      <RightNavComponent />
      <LeftNavComponent />
    </div>
  );
}
