// https://docs.cypress.io/api/introduction/api.html

describe('Home', () => {
  it('Should access home page', () => {
    cy.intercept('GET', 'http://localhost:8080/api/todo/list', { fixture: 'todolist.json' }).as('getTodoList')
    cy.visit('/')
    cy.get('.home').should('be.visible')
  })

  describe('Todo List Functioniality', function () {
    it('Should todo list visible', () => {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', { fixture: 'todolist.json' }).as('getTodoList')
      cy.visit('/')
      cy.get('#todo-list-comp').should('be.visible')
    })

    it('should todo list size correctly', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', { fixture: 'todolist.json' }).as('getTodoList')
      cy.visit('/')
      cy.wait('@getTodoList')

      cy.get('li').should('have.length', 2)
    })

    it('should empty text visible when getTodoList return empty array', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.visit('/')
      cy.wait('@getTodoList')

      cy.get('#todo-list').should('not.exist')
      cy.get('.empty-message').should('be.visible')
    })

    it('should completed icon appear on completed todo instead of complete button', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', { fixture: 'todolist_completed.json' }).as('getTodoList')
      cy.visit('/')
      cy.wait('@getTodoList')

      cy.get('li').should('exist')
        .children('.todo-check-container').should('exist')
        .children('#complete-btn').should('not.exist')
    })

    it('should complete button appear if todo is not completed', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', { fixture: 'todolist_not_completed.json' }).as('getTodoList')
      cy.visit('/')
      cy.wait('@getTodoList')

      cy.get('li').should('exist')
        .children('.todo-check-container').should('not.exist')
      cy.get('li').should('exist')
        .children('#complete-btn').should('not.exist')
    })
  })

  describe('Todo Create Functioniality', function () {
    it('should create container visible', function () {
      cy.visit('/')
      cy.get('.create-update-wrapper').should('exist')
    })

    it('should new button will show create form', function () {
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-form').should('exist')
      cy.get('#submit-button').should('exist')
    })

    it('should form submition send correct data', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.intercept('POST', 'http://localhost:8080/api/todo/create').as('createTodo')
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-title').type('Test title')
      cy.get('#todo-description').type('test description')
      cy.get('#submit-button').click()
      cy.wait('@createTodo').then((intercept) => {
        // @ts-ignore
        expect(intercept.request.body.title).to.equal('Test title')
        // @ts-ignore
        expect(intercept.request.body.description).to.equal('test description')
      })
    })

    it('should title not be empty', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.intercept('POST', 'http://localhost:8080/api/todo/create').as('createTodo')
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-description').type('test description')
      cy.get('#submit-button').click()
      cy.get('.error-container').should('exist')
    })

    it('should validation messages not be multiplexed', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.intercept('POST', 'http://localhost:8080/api/todo/create').as('createTodo')
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-description').type('test description')
      cy.get('#submit-button').click()
      cy.get('#submit-button').click()
      cy.get('#submit-button').click()
      cy.get('.messages').children().should('have.length', 1)
    })

    it('should form is empty after save transaction', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.intercept('POST', 'http://localhost:8080/api/todo/create', []).as('createTodo')
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-title').type('Test title')
      cy.get('#todo-description').type('test description')
      cy.get('#submit-button').click()
      cy.wait('@createTodo')
      cy.get('#show-container-btn').click()
      cy.get('#todo-title').should('not.have.value')
    })

    it('should the list be updated after registration.', function () {
      cy.intercept('GET', 'http://localhost:8080/api/todo/list', []).as('getTodoList')
      cy.intercept('POST', 'http://localhost:8080/api/todo/create', []).as('createTodo')
      cy.visit('/')
      cy.get('#show-container-btn').click()
      cy.get('#todo-title').type('Test title')
      cy.get('#todo-description').type('test description')
      cy.get('#submit-button').click()
      cy.wait('@createTodo')
      cy.wait('@getTodoList').should('exist')
      cy.get('#show-container-btn').click()
      cy.get('#todo-title').should('not.have.value')
    })
  })
})
