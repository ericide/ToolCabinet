//
//  QRCodeGenerator.h
//  QRCode
//
//  Created by junlei on 15/7/2.
//  Copyright (c) 2015年 HA. All rights reserved.
//


#import <UIKit/UIKit.h>

@interface QRCodeGenerator : NSObject

+ (UIImage *)qrImageForString:(NSString *)string imageSize:(CGFloat)size;
+ (UIImage *) twoDimensionCodeImage:(UIImage *)twoDimensionCode withAvatarImage:(UIImage *)avatarImage;
@end
