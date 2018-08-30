UPDATE payments SET pay_currency="TOP" WHERE pay_currency="TPC";

UPDATE currencies SET id="top" WHERE id="tpc";
UPDATE currencies SET options=REPLACE(options, "TPC", "TOP") WHERE id = "top";

UPDATE accounts SET currency_id="top" WHERE currency_id="tpc";
UPDATE bills SET currency_id="top" WHERE currency_id="tpc";
UPDATE bills SET currencies=REPLACE(currencies, "tpc", "top");
UPDATE deposits SET currency_id="top" WHERE currency_id="tpc";
UPDATE fund_sources SET currency_id="top" WHERE currency_id="tpc";
UPDATE payment_addresses SET currency_id="top" WHERE currency_id="tpc";
UPDATE payment_address_pres SET currency_id="top" WHERE currency_id="tpc";
UPDATE proofs SET currency_id="top" WHERE currency_id="tpc";
UPDATE withdraws SET currency_id="top" WHERE currency_id="tpc";

UPDATE markets SET bid_unit="top" WHERE bid_unit="tpc";
UPDATE markets SET ask_unit="top" WHERE ask_unit="tpc";
UPDATE markets SET id=REPLACE(id, "tpc", "top");

UPDATE orders SET bid="top" WHERE bid="tpc";
UPDATE orders SET ask="top" WHERE ask="tpc";
UPDATE orders SET market_id=REPLACE(market_id, "tpc", "top");

UPDATE trades SET market_id=REPLACE(market_id, "tpc", "top");
