/**
 * testdata
 * @author wreulicke
 */
@TestAnnotation
package testdata;

import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Retention;

@Retention(RetentionPolicy.RUNTIME)
@interface TestAnnotation {}