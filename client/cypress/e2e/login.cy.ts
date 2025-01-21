describe("Login Form", () => {
  it("should login a user successfully", () => {
    cy.visit("/login");
    cy.get("#email")
      .type("testuser@example.com")
      .should("have.value", "testuser@example.com");

    cy.get("#password").type("password123").should("have.value", "password123");
    cy.get("button[type='submit']").click();
    cy.url().should("eq", Cypress.config().baseUrl + "/");
    cy.contains("Logout").should("be.visible");
  });
});
