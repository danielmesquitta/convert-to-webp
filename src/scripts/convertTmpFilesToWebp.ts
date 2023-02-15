import { promises as fs } from 'fs';
import sharp from 'sharp';

import { tmpFolderPath, uploadsFolderPath } from '~/configs';

const convertTmpFilesToWebp = async () => {
  const fileNames = await fs.readdir(tmpFolderPath);

  await Promise.all(
    fileNames.map(async (fileName) => {
      if (fileName.includes('.png')) {
        await sharp(`${tmpFolderPath}/${fileName}`)
          .webp({ lossless: true })
          .toFile(`${uploadsFolderPath}/${fileName.replace('.png', '.webp')}`);
      }
    })
  );

  console.info('Done!');
};

export default convertTmpFilesToWebp;
