# DDGODELIV

This project serves as a proof of concept for a deliveries and employee management system, showcasing the application of Domain-Driven Design (DDD) principles. Developed over a one-month period, it uses Go for the backend and React for the frontend, resulting in a well-structured and scalable application.

## Current Functionality:
User Management: Users can create, manage, and view employees (drivers) within the system, including adding, editing, and deleting employee information.
Delivery Management: Drivers can be assigned deliveries, and each delivery can be associated with specific vehicles. This allows for efficient management of deliveries and tracking of driver assignments.

## Technology Stack:

### GO:
The core application logic is implemented using Go, a versatile language known for its performance and scalability, making it suitable for a deliveries and employee management system.
JWT: JSON Web Tokens (JWT) are utilized for authentication and authorization, ensuring secure access to the system's features.
SQLX: This library facilitates interaction with the PostgreSQL database, providing a convenient and efficient way to perform database operations.
Custom Middlewares: Custom middleware is implemented to handle logging, error recovery, and other tasks, enhancing the robustness and maintainability of the application.

### REACT:
React, a popular JavaScript library, is used to create the user interface of the application. React's component-based approach enables the development of modular and reusable UI components.

## Future Development:
While the project is currently on hold, it has laid a solid foundation for further development. Future plans may include:
Expanding the system to include customer management and order processing capabilities, allowing for end-to-end management of the delivery process.
Implementing advanced features such as real-time tracking of deliveries and integration with third-party mapping services.
Optimizing the application for performance and scalability to handle larger volumes of deliveries and employees.

## Benefits:

### Learning Platform:
* This project served as a valuable learning experience, honing my skills in Go, React, DDD, and various Go libraries.
* It allowed me to explore authentication, database interaction, and custom middleware implementation in Go, deepening my understanding of these concepts.
* The project provided hands-on experience with DDD principles, such as domain modeling, bounded contexts, and aggregates.

### Technical Exploration:
* The project allowed me to explore implementing authentication with JWT, ensuring secure access to the system.
* It provided practical experience with database interaction using sqlx, enabling efficient data retrieval and manipulation.
* I gained insights into developing custom middleware, allowing for flexible and extensible application behavior.
