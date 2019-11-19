const { ApolloServer } = require("apollo-server");
const { ApolloGateway } = require("@apollo/gateway");

const accountsHost = process.env["ACCOUNTS_HOST"] || "localhost:4001";
const reviewsHost = process.env["REVIEWS_HOST"] || "localhost:4002";
const productsHost = process.env["PRODUCTS_HOST"] || "localhost:4003";
const inventoryHost = process.env["INVENTORY_HOST"] || "localhost:4004";

const gateway = new ApolloGateway({
  serviceList: [
    { name: "accounts", url: `http://${accountsHost}/graphql` },
    { name: "reviews", url: `http://${reviewsHost}/graphql` },
    { name: "products", url: `http://${productsHost}/graphql` },
    { name: "inventory", url: `http://${inventoryHost}/graphql` }
  ]
});

(async () => {
  const { schema, executor } = await gateway.load();

  const server = new ApolloServer({ schema, executor });

  server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
  });
})();
