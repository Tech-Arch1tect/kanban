describe("Registration Form", () => {
  it("should register a new user successfully", () => {
    cy.visit("/register");
    cy.get("#username").type("testuser").should("have.value", "testuser");

    cy.get("#email")
      .type("testuser@example.com")
      .should("have.value", "testuser@example.com");

    cy.get("#password").type("password123").should("have.value", "password123");
    cy.get("button[type='submit']").click();
    // should eq base url/
    cy.url().should("eq", Cypress.config().baseUrl + "/");
    cy.contains("Logout").should("be.visible");
  });
});
