class PaymentService {
    processPayment(amount) {
      console.log(`Processing payment of $${amount}...`);
      return { status: "success", amount };
    }
  }
  
  module.exports = new PaymentService();
  