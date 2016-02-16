//
//  ViewController.m
//  QRCode
//
//  Created by junlei on 15/7/2.
//  Copyright (c) 2015年 HA. All rights reserved.
//

#import "ViewController.h"
#import "QRCodeGenerator.h"

@interface ViewController ()

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    UIImageView * im = [[UIImageView alloc]initWithImage:[QRCodeGenerator qrImageForString:@"http://www.baidu.com" imageSize:100]];
    [self.view addSubview:im];
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
