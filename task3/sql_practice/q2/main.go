package q2

/*
*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

/**
BEGIN;
-- 假设 A=1, B=2, 金额=100
START TRANSACTION;

-- 只有当余额>=100 时才会成功扣款（并隐式获得该行的锁）
UPDATE accounts
SET balance = balance - 100
WHERE id = 1 AND balance >= 100;

-- 如果没扣到（受影响行数=0），说明余额不足
-- （在应用层检查 ROW_COUNT()；或用下面的控制台示意）
SET @rows := ROW_COUNT();
IF @rows = 0 THEN
  ROLLBACK;
ELSE
  UPDATE accounts SET balance = balance + 100 WHERE id = 2;
  INSERT INTO transactions (from_account_id, to_account_id, amount)
  VALUES (1, 2, 100);
  COMMIT;
END IF;
*/
