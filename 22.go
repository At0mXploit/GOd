package main

func getMonthlyPrice(tier string) int {
    switch tier {
    case "basic":
        return 100 * 100 // $100.00 in pennies
    case "premium":
        return 150 * 100 // $150.00 in pennies  
    case "enterprise":
        return 500 * 100 // $500.00 in pennies
    default:
        return 0 // invalid tier
    }
}
