import LendingPool, AToken from "aave/v2"

class AaveWithdrawal extends Module {
    
    this.preTransactions = ordered([
        AToken.approve(wallet, amount)
    ])

    createWithdraw(AToken token, address user, uint amount) {
        LendingPool pool = token.POOL();
        address underlyingAssetAddress = token.UNDERLYING_ASSET_ADDRESS();

        uint userBalance = token.balanceOf(user);
        uint allowance = token.allowance(user, wallet);

        uint withdrawAmount = allowance < userBalance ? allowance : userBalance;
		withdrawAmount = withdrawAmount < amount ? withdrawAmount : amount;

        token.transferFrom(user, wallet, withdrawAmount);
        pool.withdraw(token, withdrawAmount, user);
    }
}