CREATE TABLE cars
(
    Car_id INT NOT NULL AUTO_INCREMENT,
    RegNum VARCHAR(100) NOT NULL,
    Colour VARCHAR(100) NOT NULL,
    PRIMARY KEY (car_id)
);

INSERT INTO cars(RegNum,Colour) VALUES
('MH-12-4575', 'White'),
('KA-11-8596', 'Grey'),
('GJ-03-4444', 'Black'),
('GA-43-0001', 'Black'),
('MH-43-6785', 'Black');
    